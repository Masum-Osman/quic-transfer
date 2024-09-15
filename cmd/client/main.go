package main

import (
	"log"

	"github.com/masum-osman/quic-transfer/internal/client"
	"github.com/masum-osman/quic-transfer/internal/config"
)

func main() {
	cfg, err := config.Load("client")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	c := client.New(cfg)
	if err := c.Run(); err != nil {
		log.Fatalf("Client error: %v", err)
	}
}
