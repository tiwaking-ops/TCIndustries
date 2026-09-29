// Mission handlers for Phase 9: terminal listings (lazy regeneration),
// accept/abandon/turn-in lifecycle, and player-posted bounty contracts with
// escrow. Testbed fork; generic content only. GDD-silent numbers [PROVISIONAL]
// per the owner-approved convention; GDD-given values cited.
//
// Design notes (all flagged): no new tick — expiry/refill sweep runs lazily
// inside list/accept/turn-in (deterministic for tests); terminal proximity
// uses LIVE world position (online presence required — terminals are physical);
// delivery/sample goods are destroyed as ledger sinks (no NPC receiver exists);
// group sharing is OUT (needs group-bonus design); Assault + NPC-audience
// missions are OUT (no NPCs — blocked with reason); faction-point rewards are
// zero in MVP ("if applicable" never applies yet).
package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"swg-server/internal/creatures"
	"swg-server/internal/database"
	"swg-server/internal/economy"
)

// MissionHandler serves /api/missions/*. It holds the world handler for live
// presence/position checks and WS-adjacent lookups.
type MissionHandler struct {
	db *database.DB
	w  *WorldHandler
}

// NewMissionHandler creates the handler.
func NewMissionHandler(db *database.DB, w *WorldHandler) *MissionHandler {
	return &MissionHandler{db: db, w: w}
}

// Mission types (GDD 16.3 subset reachable without NPCs).
const (
	MissionDestroyLair = "destroy_lair"
	MissionRecon       = "recon"
	MissionDelivery    = "delivery"
	MissionSample      = "sample"
)

// Mission tuning [PROVISIONAL where GDD-silent; GDD-given cited].
const (
	MissionLogCap      = 4   // GDD 16.2.3 log cap 3–5 [ASSUMPTION], midpoint
	MissionListings    = 10  // GDD 16.2.2 typically 8–15, midpoint
	MissionRefreshSecs = 420 // GDD 16.2.2 refresh 5–10 min, midpoint 7 min
	MissionExpirySecs  = 7200 // GDD 16.2.2 expiry 2 h (given)
	MissionReconRadius = 20.0 // presence radius for recon completion [PROVISIONAL]
	ContractExpirySecs = 86400 // player contracts live 24 h [PROVISIONAL]
	BountyHunterXP     = 100   // BH XP per paid contract [PROVISIONAL]
)

// HandleMissionRoute dispatches /api/missions/* sub-paths.
func (mh *MissionHandler) HandleMissionRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/missions/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "terminals" && r.Method == "GET":
		mh.terminals(w, r, charID)
	case path == "list" && r.Method == "GET":
		mh.list(w, r, charID, r.URL.Query().Get("terminal_id"))
	case path == "accept" && r.Method == "POST":
		mh.accept(w, r)
	case path == "abandon" && r.Method == "POST":
		mh.abandon(w, r)
	case path == "turn-in" && r.Method == "POST":
		mh.turnIn(w, r)
	case path == "log" && r.Method == "GET":
		mh.missionLog(w, r, charID)
	case path == "contract/post" && r.Method == "POST":
		mh.contractPost(w, r)
	case path == "contract/list" && r.Method == "GET":
		mh.contractList(w, r, charID)
	case path == "contract/cancel" && r.Method == "POST":
		mh.contractCancel(w, r)
	case path == "lairs" && r.Method == "GET":
		mh.lairs(w, r)
	case path == "lair" && r.Method == "GET":
		mh.lair(w, r, r.URL.Query().Get("lair_id"))
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown mission route"})
	}
}

// livePos returns the live world position (presence required).
func (mh *MissionHandler) livePos(charID string) (float64, float64, string, bool) {
	cl, ok := mh.w.findClient(charID)
	if !ok {
		return 0, 0, "", false
	}
	return cl.Pos.X, cl.Pos.Z, cl.Pos.Planet, true
}

// nearTerminal reports whether live position is near a terminal of category.
func (mh *MissionHandler) nearTerminal(charID, zone, category string) bool {
	x, z, planet, ok := mh.livePos(charID)
	if !ok || planet != zone {
		return false
	}
	terms, err := mh.db.TerminalsByCategory(zone, category)
	if err != nil {
		return false
	}
	for _, t := range terms {
		dx, dz := x-t.X, z-t.Z
		if dx*dx+dz*dz <= economy.TerminalSearchRadiusM*economy.TerminalSearchRadiusM {
			return true
		}
	}
	return false
}

// sweep expires past-due missions + contracts (refunding the latter) and
// refills each terminal to the listing target. Lazy: called from list/accept.
func (mh *MissionHandler) sweep() {
	now := time.Now().Unix()
	_, _ = mh.db.ExpireMissions(now)
	for _, c := range mh.expiredContracts(now) {
		_ = mh.db.RefundContract(c.ID)
	}
	for _, cat := range []string{"combat", "crafting", "bounty"} {
		terms, _ := mh.db.TerminalsByCategory("zone-0001", cat)
		for i := range terms {
			mh.refill(&terms[i], now)
		}
	}
}

func (mh *MissionHandler) expiredContracts(now int64) []database.ContractRow {
	out, _ := mh.db.ExpiredContracts(now)
	return out
}

// refill tops a terminal to MissionListings available missions.
func (mh *MissionHandler) refill(t *database.MissionTerminal, now int64) {
	have, _ := mh.db.AvailableMissions(t.ID)
	need := MissionListings - len(have)
	if need <= 0 {
		return
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < need; i++ {
		m := mh.generate(t, rng, now)
		if m == nil {
			continue
		}
		if _, err := mh.db.CreateMission(m); err != nil {
			log.Printf("mission generate failed: %v", err)
		}
	}
}

// generate builds one mission appropriate to the terminal category.
// Rewards sit inside the GDD 12.2.1 500–10k band (provisional in-band mapping).
func (mh *MissionHandler) generate(t *database.MissionTerminal, rng *rand.Rand, now int64) *database.MissionRow {
	exp := now + MissionExpirySecs
	switch t.Category {
	case "combat":
		lairs, _ := mh.db.AllLairs()
		var live []database.LairRow
		for _, l := range lairs {
			if !l.DestroyedAt.Valid {
				live = append(live, l)
			}
		}
		if len(live) == 0 {
			return nil
		}
		l := live[rng.Intn(len(live))]
		cl := 1
		if tmpl := templateCL(l.TemplateID); tmpl > 0 {
			cl = tmpl
		}
		return &database.MissionRow{
			TerminalID: t.ID, Type: MissionDestroyLair, TargetRef: l.ID,
			RewardCredit: 1500 + 500*cl, RewardXPType: "combat", RewardXP: 200,
			ExpiresAt: exp,
		}
	case "crafting":
		if rng.Intn(2) == 0 {
			schems := []string{"basic_sidearm", "basic_plate", "structure_deed", "vendor_deed"}
			s := schems[rng.Intn(len(schems))]
			qty := 2 + rng.Intn(3) // 2–4
			return &database.MissionRow{
				TerminalID: t.ID, Type: MissionDelivery, Schematic: s, Qty: qty,
				RewardCredit: 500 * qty, RewardXPType: "crafting", RewardXP: 200,
				ExpiresAt: exp,
			}
		}
		types := []string{"ferric_metal", "conductive_alloy", "structural_polymer",
			"cultured_organic", "industrial_chemical", "fibrous_flora", "filtered_water"}
		return &database.MissionRow{
			TerminalID: t.ID, Type: MissionSample, Schematic: types[rng.Intn(len(types))],
			Qty: 10, RewardCredit: 750, RewardXPType: "crafting", RewardXP: 100,
			ExpiresAt: exp,
		}
	default:
		// Bounty terminals list player contracts (posted separately) — no
		// auto-generated NPC bounties (no NPC targets exist; blocked).
		if rng.Intn(2) == 0 {
			// Recon overflow: scouting work for anyone near a terminal.
			x := t.X + (rng.Float64()*2 - 1) * 400
			z := t.Z + (rng.Float64()*2 - 1) * 400
			return &database.MissionRow{
				TerminalID: t.ID, Type: MissionRecon,
				TargetRef: fmt.Sprintf("%.0f,%.0f", x, z),
				RewardCredit: 500, RewardXPType: "scouting", RewardXP: 50,
				ExpiresAt: exp,
			}
		}
		return nil
	}
}

// templateCL returns a template's CLMax (difficulty source for reward bands).
func templateCL(templateID string) int {
	if tmpl := creatures.TemplateByID(templateID); tmpl != nil && tmpl.CLMax > 0 {
		return tmpl.CLMax
	}
	return 1
}

func (mh *MissionHandler) terminals(w http.ResponseWriter, r *http.Request, charID string) {
	out := []database.MissionTerminal{}
	for _, cat := range []string{"combat", "crafting", "bounty"} {
		terms, _ := mh.db.TerminalsByCategory("zone-0001", cat)
		out = append(out, terms...)
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"terminals": out})
}

func (mh *MissionHandler) list(w http.ResponseWriter, r *http.Request, charID, terminalID string) {
	if terminalID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "terminal_id required"})
		return
	}
	mh.sweep()
	list, err := mh.db.AvailableMissions(terminalID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	now := time.Now().Unix()
	live := []database.MissionRow{}
	for _, m := range list {
		if m.ExpiresAt > now {
			live = append(live, m)
		}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"missions": live})
}

func (mh *MissionHandler) accept(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, missionID := strOf(body, "character_id"), strOf(body, "mission_id")
	m, err := mh.db.GetMission(missionID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "mission unknown"})
		return
	}
	if m.Status != "available" || m.ExpiresAt <= time.Now().Unix() {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "mission unavailable"})
		return
	}
	log, err := mh.db.MissionLog(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "log lookup failed"})
		return
	}
	if len(log) >= MissionLogCap {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "mission log full"})
		return
	}
	if err := mh.db.SetMissionStatus(missionID, "accepted", charID); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "accept failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "accepted"})
}

func (mh *MissionHandler) abandon(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, missionID := strOf(body, "character_id"), strOf(body, "mission_id")
	m, err := mh.db.GetMission(missionID)
	if err != nil || m.AcceptedBy != charID || m.Status != "accepted" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not your active mission"})
		return
	}
	_ = mh.db.SetMissionStatus(missionID, "abandoned", charID)
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "abandoned"})
}

// turnIn validates the objective and pays reward (credits faucet + typed XP).
func (mh *MissionHandler) turnIn(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, missionID := strOf(body, "character_id"), strOf(body, "mission_id")
	m, err := mh.db.GetMission(missionID)
	if err != nil || m.AcceptedBy != charID || m.Status != "accepted" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not your active mission"})
		return
	}
	if m.ExpiresAt <= time.Now().Unix() {
		_ = mh.db.SetMissionStatus(missionID, "expired", charID)
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "mission expired"})
		return
	}
	mult := 1.0
	switch m.Type {
	case MissionDestroyLair:
		lair, err := mh.db.GetLair(m.TargetRef)
		if err != nil || !lair.DestroyedAt.Valid {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target intact"})
			return
		}
		// GDD 16.5: destroyed-by-another still pays 50% (accepter credit is
		// the full case, via the damage log).
		attackers, _ := mh.db.LairAttackers(m.TargetRef, 0)
		credited := false
		for _, a := range attackers {
			if a == charID {
				credited = true
			}
		}
		if !credited {
			mult = 0.5
		}
	case MissionRecon:
		parts := strings.Split(m.TargetRef, ",")
		if len(parts) != 2 {
			writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "bad recon target"})
			return
		}
		tx, _ := strconv.ParseFloat(parts[0], 64)
		tz, _ := strconv.ParseFloat(parts[1], 64)
		x, z, _, ok := mh.livePos(charID)
		if !ok {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "must be in world"})
			return
		}
		dx, dz := x-tx, z-tz
		if dx*dx+dz*dz > MissionReconRadius*MissionReconRadius {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target location not visited"})
			return
		}
	case MissionDelivery:
		if !mh.consumeItems(charID, m.Schematic, m.Qty) {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "goods missing"})
			return
		}
		_ = mh.db.RecordLedger(charID, 0, "mission_delivery", "sink", missionID, "")
	case MissionSample:
		if !mh.consumeStack(charID, m.Schematic, m.Qty) {
			writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "units missing"})
			return
		}
		_ = mh.db.RecordLedger(charID, 0, "mission_sample", "sink", missionID, "")
	default:
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown mission type"})
		return
	}
	credits := int(float64(m.RewardCredit) * mult)
	_ = mh.db.AddCredits(charID, credits)
	_ = mh.db.RecordLedger(charID, credits, "mission_reward", "faucet", missionID, "")
	if m.RewardXPType != "" && m.RewardXP > 0 {
		_ = mh.db.AddCharacterXP(charID, m.RewardXPType, int(float64(m.RewardXP)*mult))
	}
	_ = mh.db.SetMissionStatus(missionID, "completed", charID)
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"status": "completed", "credits": credits,
	})
}

func (mh *MissionHandler) missionLog(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	list, err := mh.db.MissionLog(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.MissionRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"missions": list})
}

// consumeItems destroys qty unequipped owned items of a schematic (delivery
// sink — no NPC receiver exists). Listed/equipped items are excluded.
func (mh *MissionHandler) consumeItems(charID, schematic string, qty int) bool {
	items, err := mh.db.GetItems(charID)
	if err != nil {
		return false
	}
	var victims []string
	for i := range items {
		if items[i].Schematic == schematic && !items[i].Equipped {
			victims = append(victims, items[i].ID)
			if len(victims) >= qty {
				break
			}
		}
	}
	if len(victims) < qty {
		return false
	}
	for _, id := range victims {
		if err := mh.db.DeleteItem(charID, id); err != nil {
			return false
		}
	}
	return true
}

// consumeStack deducts resource units of a type (sample sink).
func (mh *MissionHandler) consumeStack(charID, resourceType string, qty int) bool {
	return mh.db.ConsumeStackType(charID, resourceType, qty) == nil
}

// contractPost opens a player bounty: target must be overt at posting, escrow
// deducted upfront (anti-spam cost). Poster=killer self-claim allowed
// (zero-sum, flagged).
func (mh *MissionHandler) contractPost(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID := strOf(body, "character_id")
	targetID, err := mh.resolveName(strOf(body, "target_name"))
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "target unknown"})
		return
	}
	if targetID == charID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "cannot post on yourself"})
		return
	}
	st, err := mh.db.GetStanding(targetID)
	if err != nil || !st.Overt || st.Alignment == "neutral" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target must be overt"})
		return
	}
	amount := 0
	if f, isNum := body["amount"].(float64); isNum {
		amount = int(f)
	}
	if amount <= 0 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "amount must be positive"})
		return
	}
	id, err := mh.db.PostContract(charID, targetID, amount, time.Now().Unix()+ContractExpirySecs)
	if err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "post failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]string{"contract_id": id})
}

// contractList shows open contracts (Novice BH gate — flagged visibility rule).
func (mh *MissionHandler) contractList(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	if ok, err := mh.db.HasSkillBox(charID, "bountyhunter_novice"); err != nil || !ok {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires Novice Bounty Hunter"})
		return
	}
	list, err := mh.db.OpenContracts(time.Now().Unix())
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	if list == nil {
		list = []database.ContractRow{}
	}
	type view struct {
		ID     string `json:"id"`
		Target string `json:"target"`
		Amount int    `json:"amount"`
	}
	out := []view{}
	for _, c := range list {
		out = append(out, view{ID: c.ID, Target: mh.db.CharacterName(c.TargetID), Amount: c.Amount})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"contracts": out})
}

func (mh *MissionHandler) contractCancel(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, contractID := strOf(body, "character_id"), strOf(body, "contract_id")
	c, err := mh.db.GetContract(contractID)
	if err != nil || c.PosterID != charID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "not your contract"})
		return
	}
	if c.Status != "open" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "contract not open"})
		return
	}
	if err := mh.db.RefundContract(contractID); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "cancel failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "cancelled"})
}

func (mh *MissionHandler) resolveName(name string) (string, error) {
	return mh.db.CharacterIDByName(name)
}

// lairs lists all lairs with population (test + management introspection;
// Destroy Lair targeting reads this).
func (mh *MissionHandler) lairs(w http.ResponseWriter, r *http.Request) {
	all, err := mh.db.AllLairs()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	type view struct {
		ID        string  `json:"id"`
		Template  string  `json:"template"`
		Zone      string  `json:"zone"`
		X         float64 `json:"x"`
		Z         float64 `json:"z"`
		Destroyed bool    `json:"destroyed"`
	}
	out := []view{}
	for _, l := range all {
		out = append(out, view{
			ID: l.ID, Template: l.TemplateID, Zone: l.Zone,
			X: l.PosX, Z: l.PosZ, Destroyed: l.DestroyedAt.Valid,
		})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"lairs": out})
}

// lair returns one lair with HP + living population.
func (mh *MissionHandler) lair(w http.ResponseWriter, r *http.Request, lairID string) {
	if lairID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "lair_id required"})
		return
	}
	l, err := mh.db.GetLair(lairID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "lair unknown"})
		return
	}
	pop, _ := mh.db.LivingCountForLair(lairID)
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"id": l.ID, "hp": l.LairHP, "hp_max": l.LairHPMax,
		"population": pop, "max_population": l.MaxPopulation,
		"destroyed": l.DestroyedAt.Valid,
	})
}
