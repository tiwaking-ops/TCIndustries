// Civic HTTP handlers, part 2: guilds, groups, mentorships, mail, waypoints,
// friends, and structure permissions. Same conventions as civic_http.go.
package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"swg-server/internal/civic"
	"swg-server/internal/database"
	"swg-server/internal/protocol"
)

func databaseNullString() sql.NullString { return sql.NullString{} }
func databaseNullInt() sql.NullInt64     { return sql.NullInt64{} }

// --- Guilds ---

func (ch *CivicHandler) guildFound(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	name, tag := strOf(body, "name"), strOf(body, "tag")
	if name == "" || len(name) > 40 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "guild name required (max 40)"})
		return
	}
	if _, err := ch.db.GuildOf(charID); err == nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already in a guild"})
		return
	}
	id := fmt.Sprintf("guild-%d", time.Now().UnixNano())
	if err := ch.db.CreateGuild(id, name, tag, charID, civic.GuildRegistrarCost); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "founding failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"guild_id": id})
}

func (ch *CivicHandler) guildGet(w http.ResponseWriter, r *http.Request, charID, guildID string) {
	if guildID == "" {
		g, err := ch.db.GuildOf(charID)
		if err != nil {
			writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "not in a guild"})
			return
		}
		guildID = g.ID
	}
	g, err := ch.db.GetGuild(guildID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "guild unknown"})
		return
	}
	members, _ := ch.db.GuildMembers(guildID)
	type memberView struct {
		CharacterID string `json:"character_id"`
		Name        string `json:"name"`
		Role        string `json:"role"`
	}
	mviews := []memberView{}
	for _, m := range members {
		mviews = append(mviews, memberView{
			CharacterID: m.CharacterID, Name: ch.db.CharacterName(m.CharacterID), Role: m.Role,
		})
	}
	hall := ""
	if g.HallStructureID.Valid {
		hall = g.HallStructureID.String
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"id": g.ID, "name": g.Name, "tag": g.Tag,
		"leader": ch.db.CharacterName(g.LeaderID), "treasury": g.Treasury,
		"dues_pct": g.DuesPct, "faction": g.Faction,
		"hall_structure_id": hall, "members": mviews,
	})
}

// mustGuildOfficer loads a guild and requires officer-or-leader of the actor.
func (ch *CivicHandler) mustGuildOfficer(w http.ResponseWriter, guildID, actor string) (*database.GuildRow, bool) {
	g, err := ch.db.GetGuild(guildID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "guild unknown"})
		return nil, false
	}
	role, err := ch.db.GuildMemberRole(guildID, actor)
	if err != nil || (role != civic.GuildLeader && role != civic.GuildOfficer) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild officer"})
		return nil, false
	}
	return g, true
}

func (ch *CivicHandler) mustGuildLeader(w http.ResponseWriter, guildID, actor string) (*database.GuildRow, bool) {
	g, err := ch.db.GetGuild(guildID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "guild unknown"})
		return nil, false
	}
	if g.LeaderID != actor {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild leader"})
		return nil, false
	}
	return g, true
}

func (ch *CivicHandler) guildInvite(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID := strOf(body, "character_id"), strOf(body, "guild_id")
	if _, ok := ch.mustGuildOfficer(w, guildID, charID); !ok {
		return
	}
	invitee, ok := ch.charByName(w, strOf(body, "invitee_name"))
	if !ok {
		return
	}
	if _, err := ch.db.GuildOf(invitee); err == nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "invitee already in a guild"})
		return
	}
	id := fmt.Sprintf("ginv-%d", time.Now().UnixNano())
	if err := ch.db.CreateGuildInvite(id, guildID, charID, invitee); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "invite failed"})
		return
	}
	if cl, online := ch.w.findClient(invitee); online {
		ch.w.send(cl, protocol.MsgGuildInvited, protocol.GuildInvitedMsg{
			InviteID: id, GuildID: guildID, Inviter: ch.db.CharacterName(charID),
		})
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"invite_id": id})
}

func (ch *CivicHandler) guildAccept(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	iv, err := ch.db.GetGuildInvite(strOf(body, "invite_id"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "invite unknown"})
		return
	}
	if iv.InviteeID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "invite not addressed to you"})
		return
	}
	if _, err := ch.db.GuildOf(charID); err == nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already in a guild"})
		return
	}
	if _, err := ch.db.GetGuild(iv.GuildID); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "guild dissolved"})
		return
	}
	_ = ch.db.ResolveGuildInvite(iv.ID, "accepted")
	if err := ch.db.AddGuildMember(iv.GuildID, charID, civic.GuildMember); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "join failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "joined"})
}

func (ch *CivicHandler) guildDecline(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	iv, err := ch.db.GetGuildInvite(strOf(body, "invite_id"))
	if err != nil || iv.InviteeID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "invite unknown"})
		return
	}
	_ = ch.db.ResolveGuildInvite(iv.ID, "declined")
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "declined"})
}

func (ch *CivicHandler) guildLeave(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID := strOf(body, "character_id"), strOf(body, "guild_id")
	g, err := ch.db.GetGuild(guildID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "guild unknown"})
		return
	}
	if _, err := ch.db.GuildMemberRole(guildID, charID); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not a member"})
		return
	}
	if g.LeaderID == charID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "transfer leadership first"})
		return
	}
	_ = ch.db.RemoveGuildMember(guildID, charID)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "left"})
}

func (ch *CivicHandler) guildKick(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID, target := strOf(body, "character_id"), strOf(body, "guild_id"), strOf(body, "target_id")
	g, ok := ch.mustGuildOfficer(w, guildID, charID)
	if !ok {
		return
	}
	_ = g
	if target == g.LeaderID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "cannot kick the leader"})
		return
	}
	if _, err := ch.db.GuildMemberRole(guildID, target); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not a member"})
		return
	}
	_ = ch.db.RemoveGuildMember(guildID, target)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "kicked"})
}

func (ch *CivicHandler) guildPromote(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID, target, role := strOf(body, "character_id"), strOf(body, "guild_id"), strOf(body, "target_id"), strOf(body, "role")
	if _, ok := ch.mustGuildLeader(w, guildID, charID); !ok {
		return
	}
	if role != civic.GuildOfficer && role != civic.GuildMember {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "role must be officer or member"})
		return
	}
	if _, err := ch.db.GuildMemberRole(guildID, target); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not a member"})
		return
	}
	_ = ch.db.SetGuildMemberRole(guildID, target, role)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "role set"})
}

func (ch *CivicHandler) guildTransfer(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID, target := strOf(body, "character_id"), strOf(body, "guild_id"), strOf(body, "target_id")
	if _, ok := ch.mustGuildLeader(w, guildID, charID); !ok {
		return
	}
	if _, err := ch.db.GuildMemberRole(guildID, target); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not a member"})
		return
	}
	if err := ch.db.SetGuildLeader(guildID, target); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "transfer failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "leadership transferred"})
}

func (ch *CivicHandler) guildDisband(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID := strOf(body, "character_id"), strOf(body, "guild_id")
	if _, ok := ch.mustGuildLeader(w, guildID, charID); !ok {
		return
	}
	_ = ch.db.DisbandGuild(guildID)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "disbanded"})
}

func (ch *CivicHandler) guildSetDues(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID := strOf(body, "character_id"), strOf(body, "guild_id")
	if _, ok := ch.mustGuildLeader(w, guildID, charID); !ok {
		return
	}
	pct := intOf(body, "pct")
	if pct < 0 || pct > civic.MaxDuesPct {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "dues out of bounds"})
		return
	}
	_ = ch.db.SetGuildDues(guildID, pct)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "dues set"})
}

func (ch *CivicHandler) guildFund(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID := strOf(body, "character_id"), strOf(body, "guild_id")
	if _, err := ch.db.GuildMemberRole(guildID, charID); err != nil {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "membership required"})
		return
	}
	credits := intOf(body, "credits")
	if credits <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "credits must be positive"})
		return
	}
	if err := ch.db.FundGuildTreasury(guildID, charID, credits); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "funding failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "funded"})
}

func (ch *CivicHandler) guildInvites(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	ivs, err := ch.db.PendingGuildInvites(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if ivs == nil {
		ivs = []database.GuildInviteRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"invites": ivs})
}

// --- Groups ---

func (ch *CivicHandler) groupInvite(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	invitee, ok := ch.charByName(w, strOf(body, "invitee_name"))
	if !ok {
		return
	}
	if invitee == charID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot invite yourself"})
		return
	}
	g, err := ch.db.GroupOf(charID)
	if err != nil {
		// Inviter groupless: inviting founds the group (GDD 19.2.1).
		gid := fmt.Sprintf("group-%d", time.Now().UnixNano())
		if err := ch.db.CreateGroup(gid, charID); err != nil {
			writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "group founding failed"})
			return
		}
		g, _ = ch.db.GetGroup(gid)
	} else if g.LeaderID != charID {
		// Any member may invite (GDD 19.2.1 first sentence); the invitee joins
		// the inviter's group on accept.
		_ = g
	}
	members, _ := ch.db.GroupMembers(g.ID)
	if len(members) >= civic.GroupCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "group full"})
		return
	}
	id := fmt.Sprintf("grinv-%d", time.Now().UnixNano())
	if err := ch.db.CreateGroupInvite(id, g.ID, charID, invitee); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "invite failed"})
		return
	}
	if cl, online := ch.w.findClient(invitee); online {
		ch.w.send(cl, protocol.MsgGroupInvited, protocol.GroupInvitedMsg{
			InviteID: id, GroupID: g.ID, Inviter: ch.db.CharacterName(charID),
		})
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"invite_id": id, "group_id": g.ID})
}

func (ch *CivicHandler) groupAccept(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	iv, err := ch.db.GetGroupInvite(strOf(body, "invite_id"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "invite unknown"})
		return
	}
	if iv.InviteeID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "invite not addressed to you"})
		return
	}
	if _, err := ch.db.GroupOf(charID); err == nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already in a group"})
		return
	}
	g, err := ch.db.GetGroup(iv.GroupID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "group disbanded"})
		return
	}
	members, _ := ch.db.GroupMembers(g.ID)
	if len(members) >= civic.GroupCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "group full"})
		return
	}
	_ = ch.db.ResolveGroupInvite(iv.ID, "accepted")
	if err := ch.db.AddGroupMember(g.ID, charID); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "join failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "joined", "group_id": g.ID})
}

func (ch *CivicHandler) groupDecline(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	iv, err := ch.db.GetGroupInvite(strOf(body, "invite_id"))
	if err != nil || iv.InviteeID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "invite unknown"})
		return
	}
	_ = ch.db.ResolveGroupInvite(iv.ID, "declined")
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "declined"})
}

// leaveGroup applies leave/kick-removal with succession: leader departure
// transfers to the longest-tenured remaining member (GDD 19.5); last member
// out disbands.
func (ch *CivicHandler) leaveGroup(groupID, charID string) error {
	g, err := ch.db.GetGroup(groupID)
	if err != nil {
		return err
	}
	if err := ch.db.RemoveGroupMember(groupID, charID); err != nil {
		return err
	}
	members, _ := ch.db.GroupMembers(groupID)
	if len(members) == 0 {
		return ch.db.DisbandGroup(groupID)
	}
	if g.LeaderID == charID {
		return ch.db.SetGroupLeader(groupID, members[0])
	}
	return nil
}

func (ch *CivicHandler) groupLeave(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	g, err := ch.db.GroupOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not in a group"})
		return
	}
	if err := ch.leaveGroup(g.ID, charID); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "leave failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "left"})
}

func (ch *CivicHandler) groupKick(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, target := strOf(body, "character_id"), strOf(body, "target_id")
	g, err := ch.db.GroupOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not in a group"})
		return
	}
	if g.LeaderID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires group leader"})
		return
	}
	if target == charID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "use leave instead"})
		return
	}
	if tg, err := ch.db.GroupOf(target); err != nil || tg.ID != g.ID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not in your group"})
		return
	}
	if err := ch.leaveGroup(g.ID, target); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "kick failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "kicked"})
}

func (ch *CivicHandler) groupLootRule(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, rule := strOf(body, "character_id"), strOf(body, "rule")
	if !civic.ValidLootRule(rule) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown loot rule"})
		return
	}
	g, err := ch.db.GroupOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not in a group"})
		return
	}
	if g.LeaderID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires group leader"})
		return
	}
	_ = ch.db.SetLootRule(g.ID, rule)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "loot rule set"})
}

func (ch *CivicHandler) groupGet(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	g, err := ch.db.GroupOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusOK, map[string]interface{}{"group": nil})
		return
	}
	members, _ := ch.db.GroupMembers(g.ID)
	names := []string{}
	for _, id := range members {
		names = append(names, ch.db.CharacterName(id))
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"group": map[string]interface{}{
			"id": g.ID, "leader": ch.db.CharacterName(g.LeaderID),
			"loot_rule": g.LootRule, "members": names,
		},
	})
}

func (ch *CivicHandler) groupInvites(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	ivs, err := ch.db.PendingGroupInvites(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if ivs == nil {
		ivs = []database.GroupInviteRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"invites": ivs})
}

// --- Mentorship ---

func (ch *CivicHandler) mentorBond(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	mentorID := strOf(body, "character_id")
	protegeID, ok := ch.charByName(w, strOf(body, "protege_name"))
	if !ok {
		return
	}
	if protegeID == mentorID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot mentor yourself"})
		return
	}
	n, err := ch.db.CountSkillBoxes(protegeID)
	if err != nil || n >= civic.MentorshipMaxBoxes {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "protege above newcomer threshold"})
		return
	}
	if err := ch.db.AddBond(mentorID, protegeID,
		time.Now().AddDate(0, 0, civic.MentorshipDays).Unix()); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "bond failed: " + err.Error()})
		return
	}
	_ = ch.db.AddCharacterXP(mentorID, "mentoring", civic.MentorshipXP)
	_ = ch.db.AddCharacterXP(protegeID, "mentoring", civic.MentorshipXP)
	writeCraftJSON(w, http.StatusCreated, map[string]string{"status": "bonded"})
}

func (ch *CivicHandler) mentorList(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	bonds, err := ch.db.BondsOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if bonds == nil {
		bonds = []database.BondRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"bonds": bonds})
}

// --- Mail ---

func (ch *CivicHandler) mailSend(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	senderID := strOf(body, "character_id")
	recipientID, ok := ch.charByName(w, strOf(body, "recipient_name"))
	if !ok {
		return
	}
	subject, text := strOf(body, "subject"), strOf(body, "body")
	if len(subject) > 120 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "subject too long"})
		return
	}
	credits := intOf(body, "credits")
	var items []string
	if raw, present := body["item_ids"].([]interface{}); present {
		for _, v := range raw {
			if s, isStr := v.(string); isStr {
				items = append(items, s)
			}
		}
	}
	id := fmt.Sprintf("mail-%d", time.Now().UnixNano())
	zone := ""
	if c, err := ch.db.GetCharacterByID(senderID); err == nil {
		zone = c.Planet
	}
	if err := ch.db.SendMail(id, senderID, recipientID, subject, text,
		credits, items, zone, civic.MailCap); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "send failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"mail_id": id})
}

func (ch *CivicHandler) mailInbox(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	list, err := ch.db.Inbox(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.MailRow{}
	}
	type view struct {
		ID       string `json:"id"`
		Sender   string `json:"sender"`
		Subject  string `json:"subject"`
		Credits  int    `json:"credits_attached"`
		Claimed  bool   `json:"credits_claimed"`
		Read     bool   `json:"read"`
		HasItems bool   `json:"has_items"`
	}
	out := []view{}
	for _, m := range list {
		full, _ := ch.db.GetMail(m.ID)
		hasItems := full != nil && len(full.AttachedItems) > 0
		out = append(out, view{ID: m.ID, Sender: ch.db.CharacterName(m.SenderID),
			Subject: m.Subject, Credits: m.CreditsAttached,
			Claimed: m.CreditsClaimed, Read: m.Read, HasItems: hasItems})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"mail": out})
}

func (ch *CivicHandler) mailRead(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, mailID := strOf(body, "character_id"), strOf(body, "mail_id")
	m, err := ch.db.GetMail(mailID)
	if err != nil || m.RecipientID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "mail unknown"})
		return
	}
	_ = ch.db.MarkMailRead(mailID, charID)
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"mail": m})
}

func (ch *CivicHandler) mailClaim(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, mailID := strOf(body, "character_id"), strOf(body, "mail_id")
	zone := ""
	if c, err := ch.db.GetCharacterByID(charID); err == nil {
		zone = c.Planet
	}
	m, err := ch.db.ClaimMail(mailID, charID, zone)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "claim failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"credits": m.CreditsAttached, "items": m.AttachedItems,
	})
}

func (ch *CivicHandler) mailDelete(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	if err := ch.db.DeleteMail(strOf(body, "mail_id"), strOf(body, "character_id")); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "delete failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// --- Waypoints + friends ---

func (ch *CivicHandler) waypointAdd(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	existing, _ := ch.db.WaypointsOf(charID)
	if len(existing) >= civic.WaypointCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "waypoint list full"})
		return
	}
	zone := strOf(body, "zone")
	if zone == "" {
		if c, err := ch.db.GetCharacterByID(charID); err == nil {
			zone = c.Planet
		}
	}
	x, z := 0.0, 0.0
	if f, isNum := body["x"].(float64); isNum {
		x = f
	}
	if f, isNum := body["z"].(float64); isNum {
		z = f
	}
	source := strOf(body, "source")
	if source == "" {
		source = "manual"
	}
	id := fmt.Sprintf("wp-%d", time.Now().UnixNano())
	if err := ch.db.AddWaypoint(id, charID, strOf(body, "label"), zone, x, z, source); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "waypoint failed"})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"waypoint_id": id})
}

func (ch *CivicHandler) waypointList(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	list, err := ch.db.WaypointsOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.WaypointRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"waypoints": list})
}

// waypointShare imports a copy for an ONLINE target and pushes a chat note
// (the spoken-location-to-clickable-marker shape of GDD 18.2.4; cross-zone
// imports carry the zone with a "different planet" display rule client-side).
func (ch *CivicHandler) waypointShare(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	src, err := ch.db.GetWaypoint(strOf(body, "waypoint_id"))
	if err != nil || src.OwnerID != charID {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "waypoint unknown"})
		return
	}
	targetID, ok := ch.charByName(w, strOf(body, "target_name"))
	if !ok {
		return
	}
	target, online := ch.w.findClient(targetID)
	if !online {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "recipient not in world"})
		return
	}
	their, _ := ch.db.WaypointsOf(targetID)
	if len(their) >= civic.WaypointCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "recipient waypoint list full"})
		return
	}
	id := fmt.Sprintf("wp-%d", time.Now().UnixNano())
	if err := ch.db.AddWaypoint(id, targetID, src.Label, src.Zone, src.X, src.Z, "shared"); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "share failed"})
		return
	}
	ch.w.send(target, protocol.MsgChatMessage, protocol.ChatMessageMsg{
		SenderName: ch.db.CharacterName(charID), Channel: "system",
		Text: "shared waypoint: " + src.Label,
	})
	writeCraftJSON(w, http.StatusOK, map[string]string{"waypoint_id": id})
}

func (ch *CivicHandler) waypointDelete(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	if err := ch.db.DeleteWaypoint(strOf(body, "waypoint_id"), strOf(body, "character_id")); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "delete failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (ch *CivicHandler) friendsAdd(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	friendID, ok := ch.charByName(w, strOf(body, "friend_name"))
	if !ok {
		return
	}
	if err := ch.db.AddFriend(charID, friendID); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "favorite failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "favorited"})
}

func (ch *CivicHandler) friendsList(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	ids, err := ch.db.FriendsOf(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	type view struct {
		CharacterID string `json:"character_id"`
		Name        string `json:"name"`
		Online      bool   `json:"online"`
	}
	out := []view{}
	for _, id := range ids {
		_, online := ch.w.findClient(id)
		out = append(out, view{CharacterID: id, Name: ch.db.CharacterName(id), Online: online})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"friends": out})
}

func (ch *CivicHandler) friendsRemove(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	friendID, ok := ch.charByName(w, strOf(body, "friend_name"))
	if !ok {
		return
	}
	_ = ch.db.RemoveFriend(charID, friendID)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "removed"})
}

// --- Structure permissions ---

func (ch *CivicHandler) structurePerms(w http.ResponseWriter, r *http.Request, structureID string) {
	if structureID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "structure_id required"})
		return
	}
	p, err := ch.db.GetStructurePerms(structureID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "structure unknown"})
		return
	}
	names := func(ids []string) []string {
		out := []string{}
		for _, id := range ids {
			out = append(out, ch.db.CharacterName(id))
		}
		return out
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"entry": p.EntryPerm, "admins": names(p.Admins),
		"friends": names(p.Friends), "banned": names(p.Banned),
	})
}

// mustStructureOwner requires ownership (GDD 13.2.4: owner full control;
// admin manages but cannot redeed/transfer — permission edits stay
// owner-only).
func (ch *CivicHandler) mustStructureOwner(w http.ResponseWriter, structureID, actor string) bool {
	st, err := ch.db.GetStructure(structureID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "structure unknown"})
		return false
	}
	if st.OwnerCharacterID != actor {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires structure owner"})
		return false
	}
	return true
}

func (ch *CivicHandler) structureSetEntry(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, structureID := strOf(body, "character_id"), strOf(body, "structure_id")
	if !ch.mustStructureOwner(w, structureID, charID) {
		return
	}
	switch perm := strOf(body, "entry"); perm {
	case "public", "friends_only", "private":
		if err := ch.db.SetEntryPerm(structureID, perm); err != nil {
			writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "update failed"})
			return
		}
		writeCraftJSON(w, http.StatusOK, map[string]string{"status": "entry set"})
	default:
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "entry must be public/friends_only/private"})
	}
}

func (ch *CivicHandler) structurePermAdd(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, structureID := strOf(body, "character_id"), strOf(body, "structure_id")
	if !ch.mustStructureOwner(w, structureID, charID) {
		return
	}
	target, ok := ch.charByName(w, strOf(body, "target_name"))
	if !ok {
		return
	}
	list := strOf(body, "list")
	if err := ch.db.AddPermList(structureID, list, target); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "update failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "listed"})
}

func (ch *CivicHandler) structurePermRemove(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, structureID := strOf(body, "character_id"), strOf(body, "structure_id")
	if !ch.mustStructureOwner(w, structureID, charID) {
		return
	}
	target, ok := ch.charByName(w, strOf(body, "target_name"))
	if !ok {
		return
	}
	list := strOf(body, "list")
	if err := ch.db.RemovePermList(structureID, list, target); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "update failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "unlisted"})
}
