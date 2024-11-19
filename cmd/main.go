package main

import (
	"log"
	"log/slog"

	"github.com/supercakecrumb/go-x3ui/config"
	"github.com/supercakecrumb/go-x3ui/internal/client"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(log.Writer(), nil))

	// Initialize client
	c := client.NewClient(cfg.Server.BaseURL, cfg.Server.Username, cfg.Server.Password, cfg.Server.Insecure, logger)

	// Perform login
	logger.Info("Logging in...")
	if err := c.Login(); err != nil {
		logger.Error("Login failed", slog.String("error", err.Error()))
		return
	}
	logger.Info("Login successful!")
}