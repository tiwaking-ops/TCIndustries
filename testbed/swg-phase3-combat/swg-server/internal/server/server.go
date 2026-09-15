package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"swg-server/internal/database"
	"swg-server/internal/handlers"
)

// Server is the main game server. It wires together the database,
// HTTP API handlers, and the WebSocket world handler.
// (GDD Section 29: Server Architecture)
type Server struct {
	db       *database.DB
	auth     *handlers.AuthHandler
	char     *handlers.CharacterHandler
	world    *handlers.WorldHandler
	skills   *handlers.SkillsHandler
	craft    *handlers.CraftHandler
	econ     *handlers.EconomyHandler
	svc      *handlers.ServicesHandler
	civic    *handlers.CivicHandler
	httpAddr string
	wsAddr   string
}

// New creates and configures the server.
func New() (*Server, error) {
	dbPath := getEnv("DB_PATH", "swg.db")
	db, err := database.New(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init database: %w", err)
	}

	world := handlers.NewWorldHandler(db)
	return &Server{
		db:       db,
		auth:     handlers.NewAuthHandler(db),
		char:     handlers.NewCharacterHandler(db),
		world:    world,
		skills:   handlers.NewSkillsHandler(db),
		craft:    handlers.NewCraftHandler(db),
		econ:     handlers.NewEconomyHandler(db),
		svc:      handlers.NewServicesHandler(db),
		civic:    handlers.NewCivicHandler(db, world),
		httpAddr: getEnv("HTTP_ADDR", ":8080"),
		wsAddr:   getEnv("WS_ADDR", ":8080"),
	}, nil
}

// Run starts the HTTP server (API + WebSocket on same port for simplicity).
func (s *Server) Run() error {
	defer s.db.Close()

	// Seed trainer NPCs (GDD 7.3)
	if err := s.db.SeedTrainers(); err != nil {
		log.Printf("Warning: failed to seed trainers: %v", err)
	}

	// Phase 3: seed creature lairs/instances and start the world simulation
	// tick (creature AI, incapacitation timeouts, HAM regen).
	if err := s.world.SeedCombatWorld(); err != nil {
		log.Printf("Warning: failed to seed combat world: %v", err)
	}
	s.world.StartTickLoop()

	// Phase 4: seed resource spawns and start resource ticks (lifecycle, harvest).
	if err := s.world.SeedResourceWorld(); err != nil {
		log.Printf("Warning: failed to seed resource world: %v", err)
	}
	s.world.StartResourceTickLoop()

	// Phase 5: seed economy persistence (ledger, structures, terminal).
	if err := s.world.SeedEconomyWorld(); err != nil {
		log.Printf("Warning: failed to seed economy world: %v", err)
	}

	// Phase 6: seed service persistence (buffs).
	if err := s.world.SeedServiceWorld(); err != nil {
		log.Printf("Warning: failed to seed service world: %v", err)
	}

	// Phase 7: seed civic persistence (cities, guilds, groups, mail, social).
	if err := s.world.SeedCivicWorld(); err != nil {
		log.Printf("Warning: failed to seed civic world: %v", err)
	}

	mux := http.NewServeMux()

	// --- Public API routes (no auth required) ---
	mux.HandleFunc("POST /api/register", s.auth.Register)
	mux.HandleFunc("POST /api/login", s.auth.Login)
	mux.HandleFunc("GET /api/species", s.char.GetSpecies)

	// --- Authenticated API routes ---
	mux.HandleFunc("POST /api/characters", handlers.AuthMiddleware(s.db, s.char.CreateCharacter))
	mux.HandleFunc("GET /api/characters", handlers.AuthMiddleware(s.db, s.char.ListCharacters))
	mux.HandleFunc("GET /api/characters/{id}", handlers.AuthMiddleware(s.db, s.char.GetCharacter))

	// --- Phase 2: Skills & Professions routes ---
	mux.HandleFunc("GET /api/professions", s.skills.HandleProfessions)
	mux.HandleFunc("GET /api/trainers", s.skills.HandleTrainers)
	mux.HandleFunc("/api/characters/", handlers.AuthMiddleware(s.db, s.skills.HandleCharacterSkillsRoute))

	// --- Phase 4: Resources & Crafting routes (schematics public; rest authed) ---
	mux.HandleFunc("GET /api/craft/schematics", s.craft.HandleCraftRoute)
	mux.HandleFunc("/api/craft/", handlers.AuthMiddleware(s.db, s.craft.HandleCraftRoute))

	// --- Phase 5: Economy Infrastructure routes (all authed) ---
	mux.HandleFunc("/api/economy/", handlers.AuthMiddleware(s.db, s.econ.HandleEconomyRoute))

	// --- Phase 6: Social Support Professions routes (all authed) ---
	mux.HandleFunc("/api/services/", handlers.AuthMiddleware(s.db, s.svc.HandleServicesRoute))

	// --- Phase 7: Civic Systems routes (all authed) ---
	mux.HandleFunc("/api/civic/", handlers.AuthMiddleware(s.db, s.civic.HandleCivicRoute))

	// --- WebSocket route (auth via query param token) ---
	mux.HandleFunc("GET /ws", s.world.HandleWebSocket)

	// --- Health check ---
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("SWG Pre-CU Server starting on %s", s.httpAddr)
	log.Printf("  API:  http://localhost%s/api/*", s.httpAddr)
	log.Printf("  WS:   ws://localhost%s/ws", s.httpAddr)
	log.Printf("  DB:   %s", getEnv("DB_PATH", "swg.db"))

	if err := http.ListenAndServe(s.httpAddr, mux); err != nil {
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
