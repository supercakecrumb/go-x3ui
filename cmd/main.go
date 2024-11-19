package main

import (
	"log"
	"log/slog"

	"github.com/supercakecrumb/go-x3ui/client"
	"github.com/supercakecrumb/go-x3ui/config"
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

	inbounds, err := c.ListInbounds()
	if err != nil {
		logger.Error("Error listing inbounds:", slog.String("error", err.Error()))
		return
	}
	logger.Info("Inbounds Listed", slog.Any("Inbounds", inbounds))

	onlines, err := c.GetOnlineClients()
	if err != nil {
		logger.Error("Error getting onlines:", slog.String("error", err.Error()))
		return
	}
	logger.Info("Online clients", slog.Any("clients", onlines))

	inboundClient := c.GenerateDefaultConfig("test", 1234)
	err = c.AddClient(2, inboundClient)
	if err != nil {
		logger.Error("Creating client failed", slog.String("error", err.Error()))
		return
	}
}
