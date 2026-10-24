package handlers

import (
	"net/http"

	"tc-server/internal/database"
)

type ServicesHandler struct {
	db *database.DB
}

func NewServicesHandler(db *database.DB) *ServicesHandler {
	return &ServicesHandler{db: db}
}

func (sh *ServicesHandler) HandleServicesRoute(w http.ResponseWriter, r *http.Request) {
	charID := r.URL.Query().Get("character_id")
	if r.URL.Path == "/api/services/buffs" && r.Method == "GET" {
		sh.buffs(w, r, charID)
		return
	}
	writeError(w, http.StatusNotFound, "unknown services route")
}

func (sh *ServicesHandler) buffs(w http.ResponseWriter, r *http.Request, charID string) {
	if charID == "" {
		writeError(w, http.StatusBadRequest, "character_id required")
		return
	}
	rows, err := sh.db.ActiveBuffs(charID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "buff lookup failed")
		return
	}
	if rows == nil {
		rows = []database.BuffRow{}
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"buffs": rows})
}
