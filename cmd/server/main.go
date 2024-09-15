package main

import (
	"log"

	"github.com/masum-osman/quic-transfer/internal/config"
	"github.com/masum-osman/quic-transfer/internal/server"
)

func main() {
	cfg, err := config.Load("server")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	s := server.New(cfg)
	if err := s.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
