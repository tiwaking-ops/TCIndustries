// Civic HTTP handlers for Phase 7: city management, guilds, groups,
// mentorships, mail, waypoints, and friends. All routes authed (account
// token); the acting character travels in the body/query like the economy and
// craft routes. Testbed fork; generic content only.
package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"swg-server/internal/civic"
	"swg-server/internal/database"
)

// CivicHandler serves /api/civic/*. It holds the world handler for WS pushes
// (group/guild invite notifications, waypoint-share notes).
type CivicHandler struct {
	db *database.DB
	w  *WorldHandler
}

// NewCivicHandler creates the handler.
func NewCivicHandler(db *database.DB, w *WorldHandler) *CivicHandler {
	return &CivicHandler{db: db, w: w}
}

// HandleCivicRoute dispatches /api/civic/* sub-paths.
func (ch *CivicHandler) HandleCivicRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/civic/")
	charID := r.URL.Query().Get("character_id")
	switch {
	// Cities
	case path == "city/get" && r.Method == "GET":
		ch.cityGet(w, r, charID, r.URL.Query().Get("city_id"))
	case path == "city/list" && r.Method == "GET":
		ch.cityList(w, r)
	case path == "city/fund" && r.Method == "POST":
		ch.cityFund(w, r)
	case path == "city/set-tax" && r.Method == "POST":
		ch.citySetTax(w, r)
	case path == "city/set-zoning" && r.Method == "POST":
		ch.citySetZoning(w, r)
	case path == "city/open-election" && r.Method == "POST":
		ch.cityOpenElection(w, r)
	case path == "city/vote" && r.Method == "POST":
		ch.cityVote(w, r)
	case path == "city/leave" && r.Method == "POST":
		ch.cityLeave(w, r)
	case path == "city/evict" && r.Method == "POST":
		ch.cityEvict(w, r)
	case path == "city/ledger" && r.Method == "GET":
		ch.cityLedger(w, r, r.URL.Query().Get("city_id"))
	// Guilds
	case path == "guild/found" && r.Method == "POST":
		ch.guildFound(w, r)
	case path == "guild/get" && r.Method == "GET":
		ch.guildGet(w, r, charID, r.URL.Query().Get("guild_id"))
	case path == "guild/invite" && r.Method == "POST":
		ch.guildInvite(w, r)
	case path == "guild/accept" && r.Method == "POST":
		ch.guildAccept(w, r)
	case path == "guild/decline" && r.Method == "POST":
		ch.guildDecline(w, r)
	case path == "guild/leave" && r.Method == "POST":
		ch.guildLeave(w, r)
	case path == "guild/kick" && r.Method == "POST":
		ch.guildKick(w, r)
	case path == "guild/promote" && r.Method == "POST":
		ch.guildPromote(w, r)
	case path == "guild/transfer" && r.Method == "POST":
		ch.guildTransfer(w, r)
	case path == "guild/disband" && r.Method == "POST":
		ch.guildDisband(w, r)
	case path == "guild/set-dues" && r.Method == "POST":
		ch.guildSetDues(w, r)
	case path == "guild/fund" && r.Method == "POST":
		ch.guildFund(w, r)
	case path == "guild/invites" && r.Method == "GET":
		ch.guildInvites(w, r, charID)
	// Groups
	case path == "group/invite" && r.Method == "POST":
		ch.groupInvite(w, r)
	case path == "group/accept" && r.Method == "POST":
		ch.groupAccept(w, r)
	case path == "group/decline" && r.Method == "POST":
		ch.groupDecline(w, r)
	case path == "group/leave" && r.Method == "POST":
		ch.groupLeave(w, r)
	case path == "group/kick" && r.Method == "POST":
		ch.groupKick(w, r)
	case path == "group/loot-rule" && r.Method == "POST":
		ch.groupLootRule(w, r)
	case path == "group/get" && r.Method == "GET":
		ch.groupGet(w, r, charID)
	case path == "group/invites" && r.Method == "GET":
		ch.groupInvites(w, r, charID)
	// Mentorship
	case path == "mentor/bond" && r.Method == "POST":
		ch.mentorBond(w, r)
	case path == "mentor/list" && r.Method == "GET":
		ch.mentorList(w, r, charID)
	// Mail
	case path == "mail/send" && r.Method == "POST":
		ch.mailSend(w, r)
	case path == "mail/inbox" && r.Method == "GET":
		ch.mailInbox(w, r, charID)
	case path == "mail/read" && r.Method == "POST":
		ch.mailRead(w, r)
	case path == "mail/claim" && r.Method == "POST":
		ch.mailClaim(w, r)
	case path == "mail/delete" && r.Method == "POST":
		ch.mailDelete(w, r)
	// Waypoints + friends
	case path == "waypoint/add" && r.Method == "POST":
		ch.waypointAdd(w, r)
	case path == "waypoint/list" && r.Method == "GET":
		ch.waypointList(w, r, charID)
	case path == "waypoint/share" && r.Method == "POST":
		ch.waypointShare(w, r)
	case path == "waypoint/delete" && r.Method == "POST":
		ch.waypointDelete(w, r)
	case path == "friends/add" && r.Method == "POST":
		ch.friendsAdd(w, r)
	case path == "friends/list" && r.Method == "GET":
		ch.friendsList(w, r, charID)
	case path == "friends/remove" && r.Method == "POST":
		ch.friendsRemove(w, r)
	// Structure permissions
	case path == "structure/perms" && r.Method == "GET":
		ch.structurePerms(w, r, r.URL.Query().Get("structure_id"))
	case path == "structure/set-entry" && r.Method == "POST":
		ch.structureSetEntry(w, r)
	case path == "structure/perm-add" && r.Method == "POST":
		ch.structurePermAdd(w, r)
	case path == "structure/perm-remove" && r.Method == "POST":
		ch.structurePermRemove(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown civic route"})
	}
}

func intOf(body map[string]interface{}, key string) int {
	if f, ok := body[key].(float64); ok {
		return int(f)
	}
	return 0
}

func (ch *CivicHandler) charByName(w http.ResponseWriter, name string) (string, bool) {
	id, err := ch.db.CharacterIDByName(name)
	if err != nil || id == "" {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "character unknown"})
		return "", false
	}
	return id, true
}

// mustMayor loads a city and requires the actor to be its mayor.
func (ch *CivicHandler) mustMayor(w http.ResponseWriter, cityID, actor string) (*database.CityRow, bool) {
	c, err := ch.db.GetCity(cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "city unknown"})
		return nil, false
	}
	if !c.MayorID.Valid || c.MayorID.String != actor {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires city mayor"})
		return nil, false
	}
	return c, true
}

// --- Cities ---

func (ch *CivicHandler) cityGet(w http.ResponseWriter, r *http.Request, _, cityID string) {
	if cityID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "city_id required"})
		return
	}
	c, err := ch.db.GetCity(cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "city unknown"})
		return
	}
	citizens, _ := ch.db.CitizensOf(cityID)
	type citizenView struct {
		CharacterID string `json:"character_id"`
		Name        string `json:"name"`
		Arrears     int    `json:"arrears"`
	}
	cviews := []citizenView{}
	for _, cz := range citizens {
		cviews = append(cviews, citizenView{
			CharacterID: cz.CharacterID, Name: ch.db.CharacterName(cz.CharacterID),
			Arrears: cz.TaxArrears,
		})
	}
	n, _ := ch.db.CountStructuresInCity(c.Zone, c.CenterX, c.CenterZ, c.RadiusM)
	mayorName := ""
	if c.MayorID.Valid {
		mayorName = ch.db.CharacterName(c.MayorID.String)
	}
	election, _ := ch.db.GetOpenElection(cityID)
	electionID := ""
	var electionEnds int64
	if election != nil {
		electionID, electionEnds = election.ID, election.EndsAt
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"id": c.ID, "name": c.Name, "zone": c.Zone, "rank": c.Rank,
		"status": c.Status, "treasury": c.Treasury,
		"upkeep_weekly":         c.UpkeepWeekly,
		"upkeep_effective":      civic.EffectiveUpkeep(c.Rank, ch.w.hasActiveMayor(c, time.Now().Unix())),
		"tax_flat_weekly":       c.TaxFlatWeekly,
		"tax_vendor_pct":        c.TaxVendorPct,
		"zoning_open":           c.ZoningOpen,
		"mayor":                 mayorName,
		"election_open":         electionID,
		"election_ends":         electionEnds,
		"citizens":              cviews,
		"structures_in_radius":  n,
		"unlocks":               CityUnlocks(c.Rank),
	})
}

func (ch *CivicHandler) cityList(w http.ResponseWriter, r *http.Request) {
	all, err := ch.db.AllCities()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	type view struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Zone   string `json:"zone"`
		Rank   string `json:"rank"`
		Status string `json:"status"`
	}
	out := []view{}
	for _, c := range all {
		if c.Status == civic.CityDissolved {
			continue
		}
		out = append(out, view{ID: c.ID, Name: c.Name, Zone: c.Zone, Rank: c.Rank, Status: c.Status})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"cities": out})
}

func (ch *CivicHandler) cityFund(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, cityID := strOf(body, "character_id"), strOf(body, "city_id")
	if _, err := ch.db.GetCity(cityID); err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "city unknown"})
		return
	}
	credits := intOf(body, "credits")
	if credits <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "credits must be positive"})
		return
	}
	if err := ch.db.FundCityTreasury(cityID, charID, credits); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "funding failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "funded"})
}

func (ch *CivicHandler) citySetTax(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	c, ok := ch.mustMayor(w, strOf(body, "city_id"), strOf(body, "character_id"))
	if !ok {
		return
	}
	flat, pct := intOf(body, "flat_weekly"), intOf(body, "vendor_pct")
	if flat < 0 || flat > civic.MaxFlatTaxWeekly {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "flat tax out of bounds"})
		return
	}
	if pct < 0 || pct > civic.MaxVendorTaxPct {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "vendor tax out of bounds"})
		return
	}
	c.TaxFlatWeekly, c.TaxVendorPct = flat, pct
	if err := ch.db.UpdateCity(c); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "tax update failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "tax set"})
}

func (ch *CivicHandler) citySetZoning(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	c, ok := ch.mustMayor(w, strOf(body, "city_id"), strOf(body, "character_id"))
	if !ok {
		return
	}
	open, present := body["open"].(bool)
	if !present {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "open required"})
		return
	}
	c.ZoningOpen = open
	if err := ch.db.UpdateCity(c); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "zoning update failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "zoning set"})
}

func (ch *CivicHandler) cityOpenElection(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, cityID := strOf(body, "character_id"), strOf(body, "city_id")
	c, err := ch.db.GetCity(cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "city unknown"})
		return
	}
	if !ch.db.IsCitizen(cityID, charID) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "citizenship required"})
		return
	}
	now := time.Now().Unix()
	if now-ch.db.LastElectionEnd(cityID) < civic.Secs(civic.ElectionCooldown) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "election cooldown"})
		return
	}
	id := fmt.Sprintf("elec-%d", time.Now().UnixNano())
	if err := ch.db.OpenElection(id, cityID, now+civic.Secs(civic.ElectionPeriod)); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "open failed: " + err.Error()})
		return
	}
	_ = c
	writeCraftJSON(w, http.StatusCreated, map[string]string{"election_id": id})
}

func (ch *CivicHandler) cityVote(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, cityID, candID := strOf(body, "character_id"), strOf(body, "city_id"), strOf(body, "candidate_id")
	if !ch.db.IsCitizen(cityID, charID) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "citizenship required"})
		return
	}
	if !ch.db.IsCitizen(cityID, candID) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "candidate must be a citizen"})
		return
	}
	e, err := ch.db.GetOpenElection(cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "no open election"})
		return
	}
	if err := ch.db.CastBallot(e.ID, charID, candID); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "vote rejected (one vote per citizen)"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "ballot cast"})
}

func (ch *CivicHandler) cityLeave(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, cityID := strOf(body, "character_id"), strOf(body, "city_id")
	if !ch.db.IsCitizen(cityID, charID) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not a citizen"})
		return
	}
	_ = ch.db.RemoveCitizen(cityID, charID)
	_ = ch.db.AddOptOut(cityID, charID)
	if c, err := ch.db.GetCity(cityID); err == nil && c.MayorID.Valid && c.MayorID.String == charID {
		c.MayorID = databaseNullString()
		c.TermEnd = databaseNullInt()
		_ = ch.db.UpdateCity(c)
		ch.w.ensureElection(c, time.Now().Unix())
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "opted out"})
}

func (ch *CivicHandler) cityEvict(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	c, ok := ch.mustMayor(w, strOf(body, "city_id"), strOf(body, "character_id"))
	if !ok {
		return
	}
	target := strOf(body, "target_id")
	citizens, _ := ch.db.CitizensOf(c.ID)
	arrears := -1
	for _, cz := range citizens {
		if cz.CharacterID == target {
			arrears = cz.TaxArrears
		}
	}
	if arrears < 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target not a citizen"})
		return
	}
	if arrears <= 0 {
		// GDD 14.2.3: eviction is for non-payment of city tax.
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "eviction requires tax arrears"})
		return
	}
	_ = ch.db.RemoveCitizen(c.ID, target)
	_ = ch.db.AddOptOut(c.ID, target)
	if c.MayorID.Valid && c.MayorID.String == target {
		c.MayorID = databaseNullString()
		c.TermEnd = databaseNullInt()
		_ = ch.db.UpdateCity(c)
		ch.w.ensureElection(c, time.Now().Unix())
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "evicted"})
}

func (ch *CivicHandler) cityLedger(w http.ResponseWriter, r *http.Request, cityID string) {
	if cityID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "city_id required"})
		return
	}
	entries, err := ch.db.LedgerFor("city:" + cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "ledger lookup failed"})
		return
	}
	if entries == nil {
		entries = []database.LedgerEntry{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"entries": entries})
}
