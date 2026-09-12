package main

import (
	"log"

	"swg-server/internal/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
