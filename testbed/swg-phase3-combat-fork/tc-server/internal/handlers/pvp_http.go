// Phase 8 HTTP handlers: faction lookup, battlefield zones, and guild-war
// management. All routes are authed; the acting character travels in the
// query string (GET) or body (POST) like the other Phase 5/7 handlers.
//
// Design authority: proposals/proposal-claude-tcindustries-phase8-faction-pvp-v1.md
//   Section 7.4 (routes), D-10/D-12 (guild-war roles).
package handlers

import (
	"net/http"
	"strings"
	"time"

	"tc-server/internal/pvp"
)

// PvPHandler serves /api/pvp/* and /api/faction/*. It borrows the world
// handler's client map for presence-aware pushes.
type PvPHandler struct {
	world *WorldHandler
}

// NewPvPHandler creates the handler.
func NewPvPHandler(w *WorldHandler) *PvPHandler {
	return &PvPHandler{world: w}
}

// HandleFactionRoute dispatches /api/faction/* .
func (ph *PvPHandler) HandleFactionRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/faction/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "me" && r.Method == "GET":
		ph.factionMe(w, charID)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown faction route"})
	}
}

// HandlePvPRoute dispatches /api/pvp/* .
func (ph *PvPHandler) HandlePvPRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/pvp/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "zones" && r.Method == "GET":
		ph.pvpZones(w)
	case path == "wars" && r.Method == "GET":
		ph.pvpWars(w, charID)
	case path == "war/declare" && r.Method == "POST":
		ph.warDeclare(w, r, charID)
	case path == "war/end" && r.Method == "POST":
		ph.warEnd(w, r, charID)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown pvp route"})
	}
}

// factionMe reports the caller's faction, standing, and flag state.
func (ph *PvPHandler) factionMe(w http.ResponseWriter, charID string) {
	out := map[string]interface{}{
		"character_id": charID,
		"faction_id":   "",
		"standing":     0,
		"pvp_state":    string(pvp.FlagCovert),
	}
	db := ph.world.db
	if cf, err := db.GetCharacterFaction(charID); err == nil && cf != nil {
		out["faction_id"] = cf.FactionID
		out["standing"] = cf.Standing
		if f, err := db.GetFactionByID(cf.FactionID); err == nil && f != nil {
			out["faction_code"] = f.Code
			out["faction_name"] = f.DisplayName
		}
		if fl, err := db.GetPvPFlag(charID); err == nil && fl != nil {
			out["pvp_state"] = fl.State
			if fl.AggressionExpiresAt > 0 {
				out["aggression_expires_at"] = fl.AggressionExpiresAt
			}
		}
	}
	writeCraftJSON(w, http.StatusOK, out)
}

// pvpZones lists active battlefield zones.
func (ph *PvPHandler) pvpZones(w http.ResponseWriter) {
	db := ph.world.db
	zones, err := db.ListZones()
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "zone lookup failed"})
		return
	}
	out := make([]map[string]interface{}, 0, len(zones))
	for _, z := range zones {
		out = append(out, map[string]interface{}{
			"id": z.ID, "name": z.Name, "zone": z.Zone, "shape": z.Shape,
			"center_x": z.CenterX, "center_z": z.CenterZ, "radius_m": z.RadiusM,
			"min_x": z.MinX, "min_z": z.MinZ, "max_x": z.MaxX, "max_z": z.MaxZ,
		})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"zones": out})
}

// pvpWars lists wars involving the caller's guild.
func (ph *PvPHandler) pvpWars(w http.ResponseWriter, charID string) {
	db := ph.world.db
	g, err := db.GuildOf(charID)
	if err != nil || g == nil {
		writeCraftJSON(w, http.StatusOK, map[string]interface{}{"wars": []interface{}{}, "reason": "no_guild"})
		return
	}
	out := []map[string]interface{}{}
	wars, err := db.GuildWarsFor(g.ID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "war lookup failed"})
		return
	}
	for _, wr := range wars {
		out = append(out, map[string]interface{}{
			"id": wr.ID, "status": wr.Status,
			"declaring_guild_id": wr.DeclaringGuildID,
			"target_guild_id":    wr.TargetGuildID,
			"declared_at":        wr.DeclaredAt,
			"active_from":        wr.ActiveFrom,
		})
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"wars": out})
}

// warDeclare records a guild-war declaration (D-10 role check).
func (ph *PvPHandler) warDeclare(w http.ResponseWriter, r *http.Request, charID string) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	db := ph.world.db
	targetGuildID := strOf(body, "target_guild_id")
	if targetGuildID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "target_guild_id required"})
		return
	}
	g, err := db.GuildOf(charID)
	if err != nil || g == nil {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "not_in_guild"})
		return
	}
	if g.ID == targetGuildID {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "self_war"})
		return
	}
	// D-10: only leader/officer may declare.
	role, err := db.GuildMemberRole(g.ID, charID)
	if err != nil {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "role_lookup_failed"})
		return
	}
	allowed := false
	for _, want := range pvp.Cfg.WarDeclareRoles {
		if role == want {
			allowed = true
			break
		}
	}
	if !allowed {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild " + strings.Join(pvp.Cfg.WarDeclareRoles, " or ")})
		return
	}
	// Duplicate check: no declared/active war on the same pair.
	if db.AnyPendingOrActiveWar(g.ID, targetGuildID) {
		writeCraftJSON(w, http.StatusConflict, map[string]string{"error": "war_exists"})
		return
	}
	warID := db.NewRowID("war")
	activeFrom := time.Now().Add(pvp.Cfg.WarActivationDelay).Unix()
	if err := db.CreateGuildWar(warID, g.ID, targetGuildID, charID, activeFrom); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "declare failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusCreated, map[string]interface{}{
		"war_id": warID, "status": "declared", "active_from": activeFrom,
	})
}

// warEnd surrenders a war (D-12: either leader).
func (ph *PvPHandler) warEnd(w http.ResponseWriter, r *http.Request, charID string) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	db := ph.world.db
	warID := strOf(body, "war_id")
	if warID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "war_id required"})
		return
	}
	war, err := db.GetGuildWar(warID)
	if err != nil || war == nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "war_unknown"})
		return
	}
	if war.Status == "ended" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already_ended"})
		return
	}
	// Either guild's leader may end it (D-12).
	g, err := db.GuildOf(charID)
	if err != nil || g == nil || (g.ID != war.DeclaringGuildID && g.ID != war.TargetGuildID) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "not_a_participant"})
		return
	}
	role, err := db.GuildMemberRole(g.ID, charID)
	if err != nil || role != "leader" {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild leader"})
		return
	}
	if err := db.EndGuildWar(warID, charID, "surrender"); err != nil {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "end failed: " + err.Error()})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "ended"})
}
