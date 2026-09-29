// Services HTTP handler for Phase 6: active-buff queries. Testbed fork;
// generic content only.
package handlers

import (
	"net/http"

	"swg-server/internal/database"
)

// ServicesHandler serves /api/services/* (AuthMiddleware-wrapped).
type ServicesHandler struct {
	db *database.DB
}

// NewServicesHandler creates the handler.
func NewServicesHandler(db *database.DB) *ServicesHandler {
	return &ServicesHandler{db: db}
}

// HandleServicesRoute dispatches /api/services/* sub-paths.
func (sh *ServicesHandler) HandleServicesRoute(w http.ResponseWriter, r *http.Request) {
	charID := r.URL.Query().Get("character_id")
	if r.URL.Path == "/api/services/buffs" && r.Method == "GET" {
		sh.buffs(w, r, charID)
		return
	}
	writeCraftJSON(w, http.StatusNotFound, map[string]string{"error": "unknown services route"})
}

// buffs returns a character's active (unexpired) buffs — wall-clock expiry
// persists across logout/login (GDD 24.5).
func (sh *ServicesHandler) buffs(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeCraftJSON(w, http.StatusBadRequest, map[string]string{"error": "character_id required"})
		return
	}
	rows, err := sh.db.ActiveBuffs(charID)
	if err != nil {
		writeCraftJSON(w, http.StatusInternalServerError, map[string]string{"error": "buff lookup failed"})
		return
	}
	if rows == nil {
		rows = []database.BuffRow{}
	}
	writeCraftJSON(w, http.StatusOK, map[string]interface{}{"buffs": rows})
}
