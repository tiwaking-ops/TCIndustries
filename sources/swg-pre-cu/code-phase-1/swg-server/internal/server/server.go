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

	return &Server{
		db:       db,
		auth:     handlers.NewAuthHandler(db),
		char:     handlers.NewCharacterHandler(db),
		world:    handlers.NewWorldHandler(db),
		httpAddr: getEnv("HTTP_ADDR", ":8080"),
		wsAddr:   getEnv("WS_ADDR", ":8080"),
	}, nil
}

// Run starts the HTTP server (API + WebSocket on same port for simplicity).
func (s *Server) Run() error {
	defer s.db.Close()

	mux := http.NewServeMux()

	// --- Public API routes (no auth required) ---
	mux.HandleFunc("POST /api/register", s.auth.Register)
	mux.HandleFunc("POST /api/login", s.auth.Login)
	mux.HandleFunc("GET /api/species", s.char.GetSpecies)

	// --- Authenticated API routes ---
	mux.HandleFunc("POST /api/characters", handlers.AuthMiddleware(s.db, s.char.CreateCharacter))
	mux.HandleFunc("GET /api/characters", handlers.AuthMiddleware(s.db, s.char.ListCharacters))
	mux.HandleFunc("GET /api/characters/{id}", handlers.AuthMiddleware(s.db, s.char.GetCharacter))

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
