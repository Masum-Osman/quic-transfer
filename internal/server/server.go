package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/masum-osman/quic-transfer/internal/config"
	"github.com/masum-osman/quic-transfer/internal/protocol"
	"github.com/masum-osman/quic-transfer/internal/utils"
	"github.com/quic-go/quic-go"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() error {
	listener, err := quic.ListenAddr(
		fmt.Sprintf(":%d", s.cfg.Port),
		utils.GenerateTLSConfig(),
		&quic.Config{},
	)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer listener.Close()

	log.Printf("Server listening on :%d", s.cfg.Port)

	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn quic.Connection) {
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Printf("Failed to accept stream: %v", err)
		return
	}
	defer stream.Close()

	fileInfo, err := protocol.ReceiveFileInfo(stream)
	if err != nil {
		log.Printf("Failed to receive file info: %v", err)
		return
	}

	filePath := filepath.Join(s.cfg.FilePath, fileInfo.Name)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return
	}
	defer file.Close()

	n, err := io.Copy(file, stream)
	if err != nil {
		log.Printf("Failed to receive file: %v", err)
		return
	}

	log.Printf("Received file %s (%d bytes)", fileInfo.Name, n)
}
