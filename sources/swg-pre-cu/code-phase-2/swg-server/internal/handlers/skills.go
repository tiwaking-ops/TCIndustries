package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"swg-server/internal/database"
	"swg-server/internal/skills"
)

// SkillsHandler handles Phase 2 skill/profession API endpoints.
// GDD Sections 7, 8.2, 28.3.
type SkillsHandler struct {
	db *database.DB
}

func NewSkillsHandler(db *database.DB) *SkillsHandler {
	return &SkillsHandler{db: db}
}

// RegisterRoutes is kept for compatibility but routes are registered in server.go.
func (h *SkillsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/professions", h.HandleProfessions)
	mux.HandleFunc("/api/trainers", h.HandleTrainers)
	mux.HandleFunc("/api/characters/", h.HandleCharacterSkillsRoute)
}

// HandleProfessions returns all professions with their full skill tree data.
// GET /api/professions
func (h *SkillsHandler) HandleProfessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	professions := skills.AllProfessions()

	type treeJSON struct {
		ID    string                   `json:"tree_id"`
		Name  string                   `json:"name"`
		Boxes []map[string]interface{} `json:"boxes"`
	}
	type profJSON struct {
		ID               string                 `json:"profession_id"`
		Name             string                 `json:"name"`
		Category         string                 `json:"category"`
		XPType           string                 `json:"xp_type"`
		TotalSkillPoints int                    `json:"total_skill_points"`
		NoviceBox        map[string]interface{} `json:"novice_box"`
		MasterBox        map[string]interface{} `json:"master_box"`
		Trees            []treeJSON             `json:"trees"`
	}

	var result []profJSON
	for _, p := range professions {
		pj := profJSON{
			ID:               p.ID,
			Name:             p.Name,
			Category:         string(p.Category),
			XPType:           string(p.XPType),
			TotalSkillPoints: p.TotalSkillPoints,
			NoviceBox:        boxToJSON(p.NoviceBox),
			MasterBox:        boxToJSON(p.MasterBox),
		}
		for _, t := range p.Trees {
			tj := treeJSON{ID: t.ID, Name: t.Name}
			for _, b := range t.Boxes {
				tj.Boxes = append(tj.Boxes, boxToJSON(b))
			}
			pj.Trees = append(pj.Trees, tj)
		}
		result = append(result, pj)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func boxToJSON(b *skills.SkillBox) map[string]interface{} {
	return map[string]interface{}{
		"box_id":            b.ID,
		"profession_id":     b.ProfessionID,
		"tree_id":           b.TreeID,
		"name":              b.Name,
		"tier":              b.Tier,
		"skill_point_cost":  b.SkillPointCost,
		"xp_cost":           b.XPCost,
		"xp_type":           string(b.XPType),
		"credit_cost":       b.CreditCost,
		"granted_abilities": b.GrantedAbilities,
	}
}

// HandleTrainers returns all trainer NPCs.
// GET /api/trainers
func (h *SkillsHandler) HandleTrainers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	trainers, err := h.db.GetTrainers()
	if err != nil {
		http.Error(w, "failed to get trainers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainers)
}

// HandleCharacterSkillsRoute routes skill-related character endpoints.
// GET  /api/characters/{id}/skills   — list owned skills + XP + skill points
// POST /api/characters/{id}/train     — train a skill box
// POST /api/characters/{id}/drop-skill — drop a skill box
// POST /api/characters/{id}/earn-xp   — earn XP (placeholder action)
func (h *SkillsHandler) HandleCharacterSkillsRoute(w http.ResponseWriter, r *http.Request) {
	// Extract character ID and action from path
	// Path format: /api/characters/{id}/{action}
	path := r.URL.Path
	// Remove /api/characters/ prefix
	rest := path[len("/api/characters/"):]
	// Split into character_id and action
	parts := splitPath(rest)
	if len(parts) < 2 {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	characterID := parts[0]
	action := parts[1]

	// Verify auth
	accountID, err := getAccountIDFromRequest(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify character ownership
	char, err := h.db.GetCharacterByID(characterID)
	if err != nil {
		http.Error(w, "character not found", http.StatusNotFound)
		return
	}
	if char.AccountID != accountID {
		http.Error(w, "character does not belong to this account", http.StatusForbidden)
		return
	}

	switch action {
	case "skills":
		h.handleGetSkills(w, r, characterID)
	case "train":
		h.handleTrain(w, r, characterID)
	case "drop-skill":
		h.handleDropSkill(w, r, characterID)
	case "earn-xp":
		h.handleEarnXP(w, r, characterID)
	default:
		http.Error(w, "unknown action: "+action, http.StatusBadRequest)
	}
}

// handleGetSkills returns a character's owned skill boxes, XP pools, and available skill points.
func (h *SkillsHandler) handleGetSkills(w http.ResponseWriter, r *http.Request, characterID string) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ownedBoxes, err := h.db.GetCharacterSkillBoxes(characterID)
	if err != nil {
		http.Error(w, "failed to get skills", http.StatusInternalServerError)
		return
	}

	xpPools, err := h.db.GetCharacterXP(characterID)
	if err != nil {
		http.Error(w, "failed to get XP", http.StatusInternalServerError)
		return
	}

	credits, err := h.db.GetCharacterCredits(characterID)
	if err != nil {
		http.Error(w, "failed to get credits", http.StatusInternalServerError)
		return
	}

	spUsed := skills.TotalSkillPointsUsed(ownedBoxes)

	// Build owned box list with details
	ownedList := []map[string]interface{}{}
	for boxID := range ownedBoxes {
		if box := skills.GetSkillBox(boxID); box != nil {
			ownedList = append(ownedList, boxToJSON(box))
		}
	}

	response := map[string]interface{}{
		"character_id":           characterID,
		"owned_skills":           ownedList,
		"xp_pools":               xpPools,
		"credits":                credits,
		"skill_points_used":      spUsed,
		"skill_points_available": skills.MaxSkillPoints - spUsed,
		"skill_points_max":       skills.MaxSkillPoints,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleTrain trains a skill box for a character.
// POST /api/characters/{id}/train
// Body: {"trainer_id": "trainer_marksman_mos", "skill_box_id": "marksman_novice"}
// GDD Section 7.3: trainer checks XP, credits, prerequisites, skill points.
func (h *SkillsHandler) handleTrain(w http.ResponseWriter, r *http.Request, characterID string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		TrainerID  string `json:"trainer_id"`
		SkillBoxID string `json:"skill_box_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate skill box exists
	box := skills.GetSkillBox(req.SkillBoxID)
	if box == nil {
		http.Error(w, "skill box not found", http.StatusBadRequest)
		return
	}

	// Validate trainer exists and teaches this profession
	trainer, err := h.db.GetTrainer(req.TrainerID)
	if err != nil {
		http.Error(w, "trainer not found", http.StatusBadRequest)
		return
	}
	if trainer.ProfessionID != box.ProfessionID {
		http.Error(w, "trainer does not teach this profession", http.StatusBadRequest)
		return
	}

	// Get current owned boxes
	ownedBoxes, err := h.db.GetCharacterSkillBoxes(characterID)
	if err != nil {
		http.Error(w, "failed to get owned skills", http.StatusInternalServerError)
		return
	}

	// Check if already owned
	if ownedBoxes[req.SkillBoxID] {
		http.Error(w, "skill box already owned", http.StatusBadRequest)
		return
	}

	// Check prerequisites (GDD 6.3.3, 7.3)
	canTrain, reason := skills.PrerequisiteCheck(req.SkillBoxID, ownedBoxes)
	if !canTrain {
		http.Error(w, reason, http.StatusBadRequest)
		return
	}

	// Check XP (GDD 7.3)
	if box.XPCost > 0 {
		xpPools, err := h.db.GetCharacterXP(characterID)
		if err != nil {
			http.Error(w, "failed to get XP", http.StatusInternalServerError)
			return
		}
		currentXP := xpPools[string(box.XPType)]
		if currentXP < box.XPCost {
			http.Error(w, "insufficient "+string(box.XPType)+" XP (have "+strconv.Itoa(currentXP)+", need "+strconv.Itoa(box.XPCost)+")", http.StatusBadRequest)
			return
		}
	}

	// Check credits (GDD 7.3)
	if box.CreditCost > 0 {
		credits, err := h.db.GetCharacterCredits(characterID)
		if err != nil {
			http.Error(w, "failed to get credits", http.StatusInternalServerError)
			return
		}
		if credits < box.CreditCost {
			http.Error(w, "insufficient credits (have "+strconv.Itoa(credits)+", need "+strconv.Itoa(box.CreditCost)+")", http.StatusBadRequest)
			return
		}
	}

	// Check skill points (GDD 6.3.1)
	spUsed := skills.TotalSkillPointsUsed(ownedBoxes)
	if spUsed+box.SkillPointCost > skills.MaxSkillPoints {
		http.Error(w, "insufficient skill points (have "+strconv.Itoa(skills.MaxSkillPoints-spUsed)+" available, need "+strconv.Itoa(box.SkillPointCost)+")", http.StatusBadRequest)
		return
	}

	// All checks passed — execute training (transactional)
	// GDD 29.5: transactional integrity

	// Deduct XP
	if box.XPCost > 0 {
		if err := h.db.DeductXP(characterID, string(box.XPType), box.XPCost); err != nil {
			http.Error(w, "failed to deduct XP: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Deduct credits
	if box.CreditCost > 0 {
		if err := h.db.DeductCredits(characterID, box.CreditCost); err != nil {
			http.Error(w, "failed to deduct credits: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Insert skill box ownership
	if err := h.db.TrainSkillBox(characterID, req.SkillBoxID); err != nil {
		http.Error(w, "failed to train skill box: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Build response
	updatedOwned, _ := h.db.GetCharacterSkillBoxes(characterID)
	newSPUsed := skills.TotalSkillPointsUsed(updatedOwned)

	response := map[string]interface{}{
		"success":                true,
		"skill_box_id":           req.SkillBoxID,
		"skill_box_name":         box.Name,
		"skill_points_used":      newSPUsed,
		"skill_points_available": skills.MaxSkillPoints - newSPUsed,
		"granted_abilities":      box.GrantedAbilities,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// handleDropSkill drops a skill box and refunds skill points.
// POST /api/characters/{id}/drop-skill
// Body: {"skill_box_id": "marksman_ranged_accuracy_01"}
// GDD Section 7.4.1: dropping skills refunds skill points (not XP).
func (h *SkillsHandler) handleDropSkill(w http.ResponseWriter, r *http.Request, characterID string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		SkillBoxID string `json:"skill_box_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	box := skills.GetSkillBox(req.SkillBoxID)
	if box == nil {
		http.Error(w, "skill box not found", http.StatusBadRequest)
		return
	}

	// Check ownership
	owned, err := h.db.HasSkillBox(characterID, req.SkillBoxID)
	if err != nil || !owned {
		http.Error(w, "skill box not owned", http.StatusBadRequest)
		return
	}

	// Find dependent boxes that must also be dropped (GDD 7.4.1)
	ownedBoxes, _ := h.db.GetCharacterSkillBoxes(characterID)
	dependents := skills.DependentBoxes(req.SkillBoxID, ownedBoxes)

	// Drop the box and all dependents
	allDropped := append([]string{req.SkillBoxID}, dependents...)
	for _, boxID := range allDropped {
		if err := h.db.DropSkillBox(characterID, boxID); err != nil {
			http.Error(w, "failed to drop skill box: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Compute new skill points
	updatedOwned, _ := h.db.GetCharacterSkillBoxes(characterID)
	newSPUsed := skills.TotalSkillPointsUsed(updatedOwned)

	response := map[string]interface{}{
		"success":                true,
		"dropped_boxes":          allDropped,
		"skill_points_used":      newSPUsed,
		"skill_points_available": skills.MaxSkillPoints - newSPUsed,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleEarnXP grants XP to a character (placeholder action for Phase 2).
// POST /api/characters/{id}/earn-xp
// Body: {"xp_type": "combat", "amount": 5000}
// GDD Section 7.2: XP earned through activities. Phase 2 uses a placeholder.
func (h *SkillsHandler) handleEarnXP(w http.ResponseWriter, r *http.Request, characterID string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		XPType string `json:"xp_type"`
		Amount int    `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "amount must be positive", http.StatusBadRequest)
		return
	}

	// Validate XP type
	validTypes := map[string]bool{
		"combat": true, "crafting": true, "scouting": true,
		"medical": true, "entertainer": true, "merchant": true,
	}
	if !validTypes[req.XPType] {
		http.Error(w, "invalid XP type", http.StatusBadRequest)
		return
	}

	if err := h.db.AddCharacterXP(characterID, req.XPType, req.Amount); err != nil {
		http.Error(w, "failed to add XP: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated XP
	xpPools, _ := h.db.GetCharacterXP(characterID)

	response := map[string]interface{}{
		"success":  true,
		"xp_type":  req.XPType,
		"earned":   req.Amount,
		"xp_pools": xpPools,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// splitPath splits a URL path segment by /.
func splitPath(path string) []string {
	var parts []string
	current := ""
	for _, c := range path {
		if c == '/' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
