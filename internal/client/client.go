package client

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

type Client struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Run() error {
	conn, err := quic.DialAddr(
		context.Background(),
		fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port),
		utils.GenerateTLSConfig(),
		&quic.Config{},
	)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.CloseWithError(0, "")

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return fmt.Errorf("failed to open stream: %w", err)
	}
	defer stream.Close()

	file, err := os.Open(c.cfg.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	filename := filepath.Base(c.cfg.FilePath)
	if c.cfg.Destination != "" {
		filename = c.cfg.Destination
	}

	if err := protocol.SendFileInfo(stream, protocol.FileInfo{
		Name: filename,
		Size: fileInfo.Size(),
	}); err != nil {
		return fmt.Errorf("failed to send file info: %w", err)
	}

	n, err := io.Copy(stream, file)
	if err != nil {
		return fmt.Errorf("failed to send file: %w", err)
	}

	log.Printf("Sent file %s (%d bytes)", filename, n)
	return nil
}
