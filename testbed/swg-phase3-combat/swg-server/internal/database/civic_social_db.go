// CivicSocialDB extends the DB with Phase 7 social persistence: mail with
// item/credit attachments, waypoints, friends, and Section 13 structure
// permission lists. Testbed fork; generic only.
package database

import (
	"database/sql"
	"fmt"
)

// EnsureCivicSocialSchema creates Phase 7 social tables if absent. Idempotent.
func (db *DB) EnsureCivicSocialSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS mail (
			id TEXT PRIMARY KEY,
			sender_id TEXT NOT NULL,
			recipient_id TEXT NOT NULL,
			subject TEXT NOT NULL DEFAULT '',
			body TEXT NOT NULL DEFAULT '',
			credits_attached INTEGER NOT NULL DEFAULT 0,
			credits_claimed INTEGER NOT NULL DEFAULT 0,
			sent_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			read INTEGER NOT NULL DEFAULT 0
		)`,
		`CREATE TABLE IF NOT EXISTS mail_items (
			mail_id TEXT NOT NULL,
			item_id TEXT NOT NULL,
			PRIMARY KEY (mail_id, item_id)
		)`,
		`CREATE TABLE IF NOT EXISTS waypoints (
			id TEXT PRIMARY KEY,
			owner_id TEXT NOT NULL,
			label TEXT NOT NULL DEFAULT '',
			zone TEXT NOT NULL,
			x REAL NOT NULL,
			z REAL NOT NULL,
			source TEXT NOT NULL DEFAULT 'manual',
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS friends (
			owner_id TEXT NOT NULL,
			friend_id TEXT NOT NULL,
			added_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (owner_id, friend_id)
		)`,
	}
	for i, s := range stmts {
		if _, err := db.conn.Exec(s); err != nil {
			return fmt.Errorf("civic social schema %d failed: %w", i+1, err)
		}
	}
	_, err := db.conn.Exec(
		`CREATE INDEX IF NOT EXISTS idx_mail_recipient ON mail(recipient_id)`)
	return err
}

// MailOwnerOf returns the in-transit owner sentinel for a mail's attachments.
// Mailing an item transfers ownership out of the sender's inventory at send
// time, so the 18.5 "house condemned before claim" edge cannot occur by
// construction (GDD 18.5).
func MailOwnerOf(mailID string) string { return "mail:" + mailID }

// --- Mail ---

// MailRow maps one mail row.
type MailRow struct {
	ID               string
	SenderID         string
	RecipientID      string
	Subject          string
	Body             string
	CreditsAttached  int
	CreditsClaimed   bool
	Read             bool
	AttachedItems    []string
}

// SendMail delivers async mail with optional credit/item attachments in ONE
// transaction: sender wallet −= credits; item ownership → mail sentinel;
// mailbox cap enforced (GDD 18.2.2 [ASSUMPTION] 50). Rollback on any failure.
func (db *DB) SendMail(id, senderID, recipientID, subject, body string,
	credits int, itemIDs []string, zone string, cap int) error {
	if len(body) > 2000 {
		return fmt.Errorf("mail body too long")
	}
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var exists int
	if err := tx.QueryRow(
		`SELECT COUNT(*) FROM characters WHERE id = ?`, recipientID).Scan(&exists); err != nil || exists == 0 {
		return fmt.Errorf("recipient unknown")
	}
	var count int
	if err := tx.QueryRow(
		`SELECT COUNT(*) FROM mail WHERE recipient_id = ?`, recipientID).Scan(&count); err != nil {
		return err
	}
	if count >= cap {
		return fmt.Errorf("recipient mailbox full")
	}
	if credits < 0 {
		return fmt.Errorf("credit attachment must not be negative")
	}
	if credits > 0 {
		var bal int
		if err := tx.QueryRow(
			`SELECT credits FROM characters WHERE id = ?`, senderID).Scan(&bal); err != nil {
			return fmt.Errorf("sender unknown: %w", err)
		}
		if bal < credits {
			return fmt.Errorf("insufficient credits")
		}
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits - ? WHERE id = ?`,
			credits, senderID); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(
		`INSERT INTO mail (id, sender_id, recipient_id, subject, body, credits_attached)
		VALUES (?, ?, ?, ?, ?, ?)`,
		id, senderID, recipientID, subject, body, credits); err != nil {
		return err
	}
	for _, itemID := range itemIDs {
		var owner string
		var equipped int
		if err := tx.QueryRow(
			`SELECT owner_character_id, equipped FROM crafted_items WHERE id = ?`,
			itemID).Scan(&owner, &equipped); err != nil {
			return fmt.Errorf("attached item missing: %w", err)
		}
		if owner != senderID {
			return fmt.Errorf("attached item not sender-held")
		}
		if equipped != 0 {
			return fmt.Errorf("attached item is equipped")
		}
		if _, err := tx.Exec(
			`UPDATE crafted_items SET owner_character_id = ? WHERE id = ?`,
			MailOwnerOf(id), itemID); err != nil {
			return err
		}
		if _, err := tx.Exec(
			`INSERT INTO mail_items (mail_id, item_id) VALUES (?, ?)`,
			id, itemID); err != nil {
			return err
		}
	}
	if credits > 0 {
		if err := RecordLedgerTx(tx, senderID, -credits,
			"mail_send", "transfer", "mail:"+id, zone); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// GetMail loads one mail with its attachment item IDs.
func (db *DB) GetMail(id string) (*MailRow, error) {
	var m MailRow
	var claimed, read int
	err := db.conn.QueryRow(
		`SELECT id, sender_id, recipient_id, subject, body, credits_attached,
			credits_claimed, read FROM mail WHERE id = ?`, id,
	).Scan(&m.ID, &m.SenderID, &m.RecipientID, &m.Subject, &m.Body,
		&m.CreditsAttached, &claimed, &read)
	if err != nil {
		return nil, err
	}
	m.CreditsClaimed = claimed != 0
	m.Read = read != 0
	rows, err := db.conn.Query(
		`SELECT item_id FROM mail_items WHERE mail_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var itemID string
		if err := rows.Scan(&itemID); err != nil {
			return nil, err
		}
		m.AttachedItems = append(m.AttachedItems, itemID)
	}
	return &m, rows.Err()
}

// Inbox lists a recipient's mail (newest last).
func (db *DB) Inbox(recipientID string) ([]MailRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, sender_id, recipient_id, subject, body, credits_attached,
			credits_claimed, read FROM mail WHERE recipient_id = ? ORDER BY rowid`,
		recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []MailRow
	for rows.Next() {
		var m MailRow
		var claimed, read int
		if err := rows.Scan(&m.ID, &m.SenderID, &m.RecipientID, &m.Subject,
			&m.Body, &m.CreditsAttached, &claimed, &read); err != nil {
			return nil, err
		}
		m.CreditsClaimed = claimed != 0
		m.Read = read != 0
		out = append(out, m)
	}
	return out, rows.Err()
}

// MarkMailRead flags a mail read (recipient only).
func (db *DB) MarkMailRead(id, recipientID string) error {
	res, err := db.conn.Exec(
		`UPDATE mail SET read = 1 WHERE id = ? AND recipient_id = ?`,
		id, recipientID)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("mail unknown")
	}
	return nil
}

// ClaimMail moves pending attachments → recipient wallet/inventory in ONE
// transaction (idempotent per leg: re-claim of a claimed leg is a no-op row,
// never a double-credit).
func (db *DB) ClaimMail(id, recipientID, zone string) (*MailRow, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var m MailRow
	var claimed int
	err = tx.QueryRow(
		`SELECT id, sender_id, recipient_id, subject, body, credits_attached,
			credits_claimed FROM mail WHERE id = ? AND recipient_id = ?`,
		id, recipientID,
	).Scan(&m.ID, &m.SenderID, &m.RecipientID, &m.Subject, &m.Body,
		&m.CreditsAttached, &claimed)
	if err != nil {
		return nil, fmt.Errorf("mail unknown: %w", err)
	}
	m.CreditsClaimed = claimed != 0
	if m.CreditsAttached > 0 && !m.CreditsClaimed {
		if _, err := tx.Exec(
			`UPDATE characters SET credits = credits + ? WHERE id = ?`,
			m.CreditsAttached, recipientID); err != nil {
			return nil, err
		}
		if _, err := tx.Exec(
			`UPDATE mail SET credits_claimed = 1, read = 1 WHERE id = ?`, id); err != nil {
			return nil, err
		}
		if err := RecordLedgerTx(tx, recipientID, m.CreditsAttached,
			"mail_claim", "transfer", "mail:"+id, zone); err != nil {
			return nil, err
		}
		m.CreditsClaimed = true
	}
	rows, err := tx.Query(`SELECT item_id FROM mail_items WHERE mail_id = ?`, id)
	if err != nil {
		return nil, err
	}
	var items []string
	for rows.Next() {
		var itemID string
		if err := rows.Scan(&itemID); err != nil {
			rows.Close()
			return nil, err
		}
		items = append(items, itemID)
	}
	rows.Close()
	for _, itemID := range items {
		var owner string
		if err := tx.QueryRow(
			`SELECT owner_character_id FROM crafted_items WHERE id = ?`,
			itemID).Scan(&owner); err != nil {
			return nil, fmt.Errorf("attached item missing: %w", err)
		}
		if owner == recipientID {
			continue // already claimed (idempotent leg)
		}
		if owner != MailOwnerOf(id) {
			return nil, fmt.Errorf("attached item no longer in transit")
		}
		if _, err := tx.Exec(
			`UPDATE crafted_items SET owner_character_id = ? WHERE id = ?`,
			recipientID, itemID); err != nil {
			return nil, err
		}
	}
	m.AttachedItems = items
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &m, nil
}

// UnclaimedAttachments reports whether a mail still holds value (delete guard:
// GDD 18.2.2 — never auto-delete with attachments unclaimed; manual delete is
// likewise refused while value is attached).
func (db *DB) UnclaimedAttachments(id string) (bool, error) {
	var credits, claimed int
	if err := db.conn.QueryRow(
		`SELECT credits_attached, credits_claimed FROM mail WHERE id = ?`,
		id).Scan(&credits, &claimed); err != nil {
		return false, err
	}
	if credits > 0 && claimed == 0 {
		return true, nil
	}
	var n int
	if err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM crafted_items WHERE owner_character_id = ?`,
		MailOwnerOf(id)).Scan(&n); err != nil {
		return false, err
	}
	return n > 0, nil
}

// DeleteMail removes a mail (recipient only; refused while value attached).
func (db *DB) DeleteMail(id, recipientID string) error {
	var rec string
	if err := db.conn.QueryRow(
		`SELECT recipient_id FROM mail WHERE id = ?`, id).Scan(&rec); err != nil {
		return fmt.Errorf("mail unknown")
	}
	if rec != recipientID {
		return fmt.Errorf("not mail recipient")
	}
	held, err := db.UnclaimedAttachments(id)
	if err != nil {
		return err
	}
	if held {
		return fmt.Errorf("claim attachments before deleting")
	}
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := tx.Exec(`DELETE FROM mail_items WHERE mail_id = ?`, id); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM mail WHERE id = ?`, id); err != nil {
		return err
	}
	return tx.Commit()
}

// --- Waypoints ---

// WaypointRow maps one waypoints row.
type WaypointRow struct {
	ID      string
	OwnerID string
	Label   string
	Zone    string
	X, Z    float64
	Source  string
}

// AddWaypoint records a waypoint (per-owner cap enforced by caller).
func (db *DB) AddWaypoint(id, ownerID, label, zone string, x, z float64, source string) error {
	_, err := db.conn.Exec(
		`INSERT INTO waypoints (id, owner_id, label, zone, x, z, source)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, id, ownerID, label, zone, x, z, source)
	return err
}

// WaypointsOf lists an owner's waypoints.
func (db *DB) WaypointsOf(ownerID string) ([]WaypointRow, error) {
	rows, err := db.conn.Query(
		`SELECT id, owner_id, label, zone, x, z, source FROM waypoints
		WHERE owner_id = ? ORDER BY rowid`, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []WaypointRow
	for rows.Next() {
		var w WaypointRow
		if err := rows.Scan(&w.ID, &w.OwnerID, &w.Label, &w.Zone, &w.X, &w.Z, &w.Source); err != nil {
			return nil, err
		}
		out = append(out, w)
	}
	return out, rows.Err()
}

// GetWaypoint loads one waypoint.
func (db *DB) GetWaypoint(id string) (*WaypointRow, error) {
	var w WaypointRow
	err := db.conn.QueryRow(
		`SELECT id, owner_id, label, zone, x, z, source FROM waypoints WHERE id = ?`,
		id).Scan(&w.ID, &w.OwnerID, &w.Label, &w.Zone, &w.X, &w.Z, &w.Source)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

// DeleteWaypoint removes an owned waypoint.
func (db *DB) DeleteWaypoint(id, ownerID string) error {
	res, err := db.conn.Exec(
		`DELETE FROM waypoints WHERE id = ? AND owner_id = ?`, id, ownerID)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("waypoint unknown")
	}
	return nil
}

// --- Friends ---

// AddFriend records a favorite (idempotent).
func (db *DB) AddFriend(ownerID, friendID string) error {
	if ownerID == friendID {
		return fmt.Errorf("cannot favorite yourself")
	}
	var exists int
	if err := db.conn.QueryRow(
		`SELECT COUNT(*) FROM characters WHERE id = ?`, friendID).Scan(&exists); err != nil || exists == 0 {
		return fmt.Errorf("character unknown")
	}
	_, err := db.conn.Exec(
		`INSERT OR IGNORE INTO friends (owner_id, friend_id) VALUES (?, ?)`,
		ownerID, friendID)
	return err
}

// RemoveFriend deletes a favorite.
func (db *DB) RemoveFriend(ownerID, friendID string) error {
	_, err := db.conn.Exec(
		`DELETE FROM friends WHERE owner_id = ? AND friend_id = ?`,
		ownerID, friendID)
	return err
}

// FriendsOf lists favorites.
func (db *DB) FriendsOf(ownerID string) ([]string, error) {
	rows, err := db.conn.Query(
		`SELECT friend_id FROM friends WHERE owner_id = ? ORDER BY added_at`, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// --- Structure permissions (GDD 13.2.4) ---

// StructurePerms bundles a structure's permission state.
type StructurePerms struct {
	EntryPerm string
	Admins    []string
	Friends   []string
	Banned    []string
}

// GetStructurePerms loads permission state.
func (db *DB) GetStructurePerms(structureID string) (*StructurePerms, error) {
	var entry sql.NullString
	if err := db.conn.QueryRow(
		`SELECT entry_perm FROM structures WHERE id = ?`, structureID).Scan(&entry); err != nil {
		return nil, fmt.Errorf("structure unknown")
	}
	p := &StructurePerms{EntryPerm: "public"}
	if entry.Valid && entry.String != "" {
		p.EntryPerm = entry.String
	}
	for _, q := range []struct {
		table string
		dst   *[]string
	}{
		{"structure_admins", &p.Admins},
		{"structure_friends", &p.Friends},
		{"structure_banned", &p.Banned},
	} {
		rows, err := db.conn.Query(
			`SELECT character_id FROM `+q.table+` WHERE structure_id = ?`, structureID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var id string
			if err := rows.Scan(&id); err != nil {
				rows.Close()
				return nil, err
			}
			*q.dst = append(*q.dst, id)
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return p, nil
}

// SetEntryPerm sets public/friends_only/private.
func (db *DB) SetEntryPerm(structureID, perm string) error {
	_, err := db.conn.Exec(
		`UPDATE structures SET entry_perm = ? WHERE id = ?`, perm, structureID)
	return err
}

// permTable maps a permission list name to its table (the banned list table
// is structure_banned — irregular plural, mapped explicitly).
func permTable(list string) (string, error) {
	switch list {
	case "admin":
		return "structure_admins", nil
	case "friend":
		return "structure_friends", nil
	case "banned":
		return "structure_banned", nil
	default:
		return "", fmt.Errorf("unknown permission list")
	}
}

// AddPermList adds a character to a permission list.
func (db *DB) AddPermList(structureID, list, charID string) error {
	table, err := permTable(list)
	if err != nil {
		return err
	}
	_, err = db.conn.Exec(
		`INSERT OR IGNORE INTO `+table+` (structure_id, character_id)
		VALUES (?, ?)`, structureID, charID)
	return err
}

// RemovePermList removes a character from a permission list.
func (db *DB) RemovePermList(structureID, list, charID string) error {
	table, err := permTable(list)
	if err != nil {
		return err
	}
	_, err = db.conn.Exec(
		`DELETE FROM `+table+` WHERE structure_id = ? AND character_id = ?`,
		structureID, charID)
	return err
}

// IsPermListed reports list membership.
func (db *DB) IsPermListed(structureID, list, charID string) bool {
	table, err := permTable(list)
	if err != nil {
		return false
	}
	var n int
	_ = db.conn.QueryRow(
		`SELECT COUNT(*) FROM `+table+` WHERE structure_id = ? AND character_id = ?`,
		structureID, charID).Scan(&n)
	return n > 0
}

// CanManageStructure reports owner-or-admin (GDD 13.2.4: owner full control;
// admin decorates/manages storage but cannot redeed/transfer — funding and
// permission edits are the managed actions in this build).
func (db *DB) CanManageStructure(structureID, charID string) bool {
	st, err := db.GetStructure(structureID)
	if err != nil {
		return false
	}
	if st.OwnerCharacterID == charID {
		return true
	}
	return db.IsPermListed(structureID, "admin", charID)
}

// MailTimestamp was a UnixNano mail-ID helper; removed in the Phase 10 B6
// portability sweep (UnixNano is not unique on Windows' 15.6 ms clock).
// Mail IDs now mint via newRowID at their call sites. No callers existed.
