// Faction HTTP handlers for Phase 8: alignment declaration, status/roster,
// guild alignment labels, base windows/info, and city PvP permission. All
// routes authed (account token); the acting character travels in the
// body/query like the civic routes. Testbed fork; generic content only.
package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"swg-server/internal/database"
	"swg-server/internal/faction"
)

// FactionHandler serves /api/faction/*. It holds the world handler for
// presence-filtered roster queries.
type FactionHandler struct {
	db *database.DB
	w  *WorldHandler
}

// NewFactionHandler creates the handler.
func NewFactionHandler(db *database.DB, w *WorldHandler) *FactionHandler {
	return &FactionHandler{db: db, w: w}
}

// HandleFactionRoute dispatches /api/faction/* sub-paths.
func (fh *FactionHandler) HandleFactionRoute(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/faction/")
	charID := r.URL.Query().Get("character_id")
	switch {
	case path == "declare" && r.Method == "POST":
		fh.declare(w, r)
	case path == "status" && r.Method == "GET":
		fh.status(w, r, charID)
	case path == "roster" && r.Method == "GET":
		fh.roster(w, r, charID)
	case path == "guild/set-alignment" && r.Method == "POST":
		fh.guildSetAlignment(w, r)
	case path == "base/info" && r.Method == "GET":
		fh.baseInfo(w, r, r.URL.Query().Get("base_id"))
	case path == "base/window" && r.Method == "POST":
		fh.baseWindow(w, r)
	case path == "city/set-pvp" && r.Method == "POST":
		fh.citySetPvP(w, r)
	default:
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown faction route"})
	}
}

// declare changes alignment. Moving TO neutral is always allowed (leaving is
// safe by §15.1 opt-in spirit); declaring a non-neutral alignment within the
// cooldown of the last change is refused. Every change resets points +
// progress and stamps the change (flagged uniform rule).
func (fh *FactionHandler) declare(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, alignment := strOf(body, "character_id"), strOf(body, "alignment")
	if !faction.ValidAlignment(alignment) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown alignment"})
		return
	}
	st, err := fh.db.GetStanding(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "standing unavailable"})
		return
	}
	if st.Alignment == alignment {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "already holds that alignment"})
		return
	}
	now := time.Now().Unix()
	if alignment != faction.AlignNeutral && st.LastChange.Valid &&
		now-st.LastChange.Int64 < faction.Secs(faction.SwitchCooldown) {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "alignment change cooling down"})
		return
	}
	st.Alignment = alignment
	st.Points = 0
	st.Overt = false // leaving or joining clears the flag; neutrals can't hold it
	st.OvertSince = sql.NullInt64{}
	st.LastChange = sql.NullInt64{Int64: now, Valid: true}
	if err := fh.db.SetStanding(st); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "declaration failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "declared"})
}

func (fh *FactionHandler) status(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	fh.w.maybeDecayPoints(charID)
	st, err := fh.db.GetStanding(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "standing unavailable"})
		return
	}
	events, _ := fh.db.PointEvents(charID)
	if events == nil {
		events = []database.PointEvent{}
	}
	recent := events
	if len(recent) > 10 {
		recent = recent[len(recent)-10:]
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"alignment": st.Alignment, "overt": st.Overt, "points": st.Points,
		"rank": faction.RankFor(st.Points), "events": recent,
	})
}

// roster lists ONLINE overt members of the caller's alignment (Phase 9 bounty
// input surface; presence-filtered here).
func (fh *FactionHandler) roster(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	st, err := fh.db.GetStanding(charID)
	if err != nil || st.Alignment == faction.AlignNeutral {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "no alignment"})
		return
	}
	ids, err := fh.db.OvertMembers(st.Alignment)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "lookup failed"})
		return
	}
	out := []string{}
	for _, id := range ids {
		if _, online := fh.w.findClient(id); online {
			out = append(out, fh.db.CharacterName(id))
		}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"overt": out})
}

func (fh *FactionHandler) guildSetAlignment(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, guildID, alignment := strOf(body, "character_id"), strOf(body, "guild_id"), strOf(body, "alignment")
	if !faction.ValidAlignment(alignment) {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "unknown alignment"})
		return
	}
	g, err := fh.db.GetGuild(guildID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "guild unknown"})
		return
	}
	if g.LeaderID != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild leader"})
		return
	}
	if err := fh.db.SetGuildAlignment(guildID, alignment); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "update failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "alignment set"})
}

func (fh *FactionHandler) baseInfo(w http.ResponseWriter, r *http.Request, baseID string) {
	if baseID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "base_id required"})
		return
	}
	b, err := fh.db.GetBase(baseID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "base unknown"})
		return
	}
	var ws, we int64
	if b.WindowStart.Valid {
		ws = b.WindowStart.Int64
	}
	if b.WindowEnd.Valid {
		we = b.WindowEnd.Int64
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{
		"id": b.ID, "guild_id": b.GuildID, "faction": b.Faction,
		"zone": b.Zone, "x": b.PosX, "z": b.PosZ,
		"hp": b.HP, "hp_max": b.HPMax, "status": b.Status,
		"window_start": ws, "window_end": we,
	})
}

// baseWindow sets (or replaces) a base's single vulnerability window
// (officer+ of the owning guild). Duration cap 2 h [PROVISIONAL, GDD-silent].
func (fh *FactionHandler) baseWindow(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, baseID := strOf(body, "character_id"), strOf(body, "base_id")
	b, err := fh.db.GetBase(baseID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "base unknown"})
		return
	}
	role, err := fh.db.GuildMemberRole(b.GuildID, charID)
	if err != nil || (role != "leader" && role != "officer") {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires guild officer"})
		return
	}
	var start, end int64
	if f, isNum := body["window_start"].(float64); isNum {
		start = int64(f)
	}
	if f, isNum := body["window_end"].(float64); isNum {
		end = int64(f)
	}
	now := time.Now().Unix()
	// Start tolerance (30 s past) covers declaration-to-request clock skew;
	// the window still must be positively bounded and ≤ 2 h long.
	if start < now-30 || end <= start || end-start > 2*60*60 {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "window out of bounds"})
		return
	}
	b.WindowStart = sql.NullInt64{Int64: start, Valid: true}
	b.WindowEnd = sql.NullInt64{Int64: end, Valid: true}
	if err := fh.db.UpdateBase(b); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "window update failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "window set"})
}

// citySetPvP toggles a city's PvP permission (mayor only; default allow).
func (fh *FactionHandler) citySetPvP(w http.ResponseWriter, r *http.Request) {
	body, ok := decodeBody(w, r)
	if !ok {
		return
	}
	charID, cityID := strOf(body, "character_id"), strOf(body, "city_id")
	c, err := fh.db.GetCity(cityID)
	if err != nil {
		writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "city unknown"})
		return
	}
	if !c.MayorID.Valid || c.MayorID.String != charID {
		writeCraftJSON(w, http.StatusForbidden, map[string]string{"error": "requires city mayor"})
		return
	}
	allowed, present := body["allowed"].(bool)
	if !present {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "allowed required"})
		return
	}
	if err := fh.db.SetCityPvP(cityID, allowed); err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "update failed"})
		return
	}
	writeCraftJSON(w, http.StatusOK, map[string]string{"status": "pvp set"})
}
