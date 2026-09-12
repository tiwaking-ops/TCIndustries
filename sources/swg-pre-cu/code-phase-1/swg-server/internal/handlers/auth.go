package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"swg-server/internal/database"
	"swg-server/internal/models"
	"swg-server/internal/protocol"
	"swg-server/internal/species"
)

type contextKey string

const accountIDKey contextKey = "account_id"

// AuthHandler handles account registration, login, and session management.
type AuthHandler struct {
	db *database.DB
}

func NewAuthHandler(db *database.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

// Register handles POST /api/register (GDD Section 27.3)
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req protocol.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	if len(req.Username) < 3 || len(req.Username) > 20 {
		writeError(w, http.StatusBadRequest, "username must be 3-20 characters")
		return
	}
	if len(req.Password) < 6 {
		writeError(w, http.StatusBadRequest, "password must be at least 6 characters")
		return
	}

	_, err := h.db.GetAccountByUsername(req.Username)
	if err == nil {
		writeError(w, http.StatusConflict, "username already exists")
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	accountID := uuid.New().String()
	if err := h.db.CreateAccount(accountID, req.Username, string(hash)); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create account")
		return
	}

	writeJSON(w, http.StatusCreated, protocol.RegisterResponse{
		AccountID: accountID,
		Username:  req.Username,
	})
}

// Login handles POST /api/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req protocol.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	account, err := h.db.GetAccountByUsername(strings.TrimSpace(req.Username))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid username or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(req.Password)); err != nil {
		writeError(w, http.StatusUnauthorized, "invalid username or password")
		return
	}

	token := generateToken(account.ID)
	writeJSON(w, http.StatusOK, protocol.LoginResponse{
		AccountID: account.ID,
		Token:     token,
		Username:  account.Username,
	})
}

func generateToken(accountID string) string {
	return fmt.Sprintf("%s:%d", accountID, time.Now().Unix())
}

func ParseToken(token string) (string, error) {
	parts := strings.SplitN(token, ":", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid token format")
	}
	return parts[0], nil
}

// --- Character Handlers ---

type CharacterHandler struct {
	db *database.DB
}

func NewCharacterHandler(db *database.DB) *CharacterHandler {
	return &CharacterHandler{db: db}
}

// GetSpecies handles GET /api/species (GDD 6.2.1)
func (h *CharacterHandler) GetSpecies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	allSpecies := species.SpeciesList()
	speciesList := make([]protocol.SpeciesInfo, 0, len(allSpecies))
	for _, sp := range allSpecies {
		speciesList = append(speciesList, protocol.SpeciesInfo{
			ID:          string(sp.ID),
			DisplayName: sp.DisplayName,
		})
	}

	writeJSON(w, http.StatusOK, protocol.SpeciesListResponse{Species: speciesList})
}

// CreateCharacter handles POST /api/characters (GDD Section 6)
func (h *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	accountID, err := getAccountIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "not authenticated")
		return
	}

	var req protocol.CreateCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) < 3 || len(req.Name) > 30 {
		writeError(w, http.StatusBadRequest, "character name must be 3-30 characters")
		return
	}

	if !species.IsValidSpecies(species.SpeciesID(req.Species)) {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("invalid species: %s", req.Species))
		return
	}

	exists, err := h.db.GetCharacterNameExists(req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	if exists {
		writeError(w, http.StatusConflict, "character name already taken")
		return
	}

	if req.Appearance.BodyType == "" {
		req.Appearance.BodyType = "average"
	}
	if req.Appearance.SkinColor == "" {
		req.Appearance.SkinColor = "default"
	}
	if req.Appearance.Height == 0 {
		req.Appearance.Height = 1.0
	}

	character := &models.Character{
		ID:        uuid.New().String(),
		AccountID: accountID,
		Name:      req.Name,
		Species:   req.Species,
		Appearance: models.CharacterAppearance{
			BodyType:  req.Appearance.BodyType,
			SkinColor: req.Appearance.SkinColor,
			HairStyle: req.Appearance.HairStyle,
			HairColor: req.Appearance.HairColor,
			FaceType:  req.Appearance.FaceType,
			EyeColor:  req.Appearance.EyeColor,
			Height:    req.Appearance.Height,
		},
		PosX:    models.DefaultSpawnX,
		PosY:    models.DefaultSpawnY,
		PosZ:    models.DefaultSpawnZ,
		Heading: 0,
		Planet:  models.DefaultPlanet,
	}

	if err := h.db.CreateCharacter(character); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create character")
		return
	}

	writeJSON(w, http.StatusCreated, protocol.CreateCharacterResponse{
		CharacterID: character.ID,
		Name:        character.Name,
		Species:     character.Species,
	})
}

// ListCharacters handles GET /api/characters
func (h *CharacterHandler) ListCharacters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	accountID, err := getAccountIDFromRequest(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "not authenticated")
		return
	}

	chars, err := h.db.GetCharactersByAccount(accountID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}

	summaries := make([]protocol.CharacterSummary, 0, len(chars))
	for _, c := range chars {
		summaries = append(summaries, protocol.CharacterSummary{
			ID:      c.ID,
			Name:    c.Name,
			Species: c.Species,
			Planet:  c.Planet,
		})
	}

	writeJSON(w, http.StatusOK, protocol.ListCharactersResponse{Characters: summaries})
}

// GetCharacter handles GET /api/characters/{id}
func (h *CharacterHandler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	charID := r.PathValue("id")
	if charID == "" {
		writeError(w, http.StatusBadRequest, "character ID required")
		return
	}

	charWithHAM, err := h.db.GetCharacterByID(charID)
	if err != nil {
		writeError(w, http.StatusNotFound, "character not found")
		return
	}

	writeJSON(w, http.StatusOK, protocol.GetCharacterResponse{
		ID:      charWithHAM.ID,
		Name:    charWithHAM.Name,
		Species: charWithHAM.Species,
		Appearance: protocol.CharacterAppearanceReq{
			BodyType:  charWithHAM.Appearance.BodyType,
			SkinColor: charWithHAM.Appearance.SkinColor,
			HairStyle: charWithHAM.Appearance.HairStyle,
			HairColor: charWithHAM.Appearance.HairColor,
			FaceType:  charWithHAM.Appearance.FaceType,
			EyeColor:  charWithHAM.Appearance.EyeColor,
			Height:    charWithHAM.Appearance.Height,
		},
		Position: protocol.Position{
			X: charWithHAM.PosX,
			Y: charWithHAM.PosY,
			Z: charWithHAM.PosZ,
		},
		Planet: charWithHAM.Planet,
		HAM: protocol.HAMState{
			Health:       charWithHAM.HAM.Health,
			Strength:     charWithHAM.HAM.Strength,
			Constitution: charWithHAM.HAM.Constitution,
			Action:       charWithHAM.HAM.Action,
			Quickness:    charWithHAM.HAM.Quickness,
			Stamina:      charWithHAM.HAM.Stamina,
			Mind:         charWithHAM.HAM.Mind,
			Focus:        charWithHAM.HAM.Focus,
			Willpower:    charWithHAM.HAM.Willpower,
		},
	})
}

// --- Middleware ---

// AuthMiddleware validates the Bearer token and injects account ID into context.
func AuthMiddleware(db *database.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeError(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			writeError(w, http.StatusUnauthorized, "invalid authorization format")
			return
		}

		accountID, err := ParseToken(parts[1])
		if err != nil {
			writeError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		_, err = db.GetAccountByID(accountID)
		if err != nil {
			writeError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), accountIDKey, accountID)
		next(w, r.WithContext(ctx))
	}
}

func getAccountIDFromRequest(r *http.Request) (string, error) {
	val := r.Context().Value(accountIDKey)
	if val == nil {
		return "", errors.New("not authenticated")
	}
	accountID, ok := val.(string)
	if !ok {
		return "", errors.New("invalid account ID in context")
	}
	return accountID, nil
}

// --- Helpers ---

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, protocol.ErrorResponse{Error: message})
}
