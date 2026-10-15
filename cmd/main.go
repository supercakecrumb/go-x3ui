package main

import (
	"log"
	"log/slog"
	"os"

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
	logger := newLogger("debug")

	// Initialize client
	c := client.NewClient(cfg.Server.BaseURL, cfg.Server.Username, cfg.Server.Password, cfg.Server.Insecure, logger)

	// // Perform login
	// logger.Info("Logging in...")
	// if err := c.Login(); err != nil {
	// 	logger.Error("Login failed", slog.String("error", err.Error()))
	// 	return
	// }
	// logger.Info("Login successful!")

	// spew.Dump(c.Resty.GetClient().Jar)

	inbounds, err := c.ListInbounds()
	if err != nil {
		logger.Error("Error listing inbounds:", slog.String("error", err.Error()))
		return
	}
	logger.Info("Inbounds Listed", slog.Int("count", len(inbounds)))

	onlines, err := c.GetOnlineClients()
	if err != nil {
		logger.Error("Error getting onlines:", slog.String("error", err.Error()))
		return
	}
	logger.Info("Online clients", slog.Int("count", len(onlines)))

	// inboundClient := c.GenerateDefaultInboundClient("testtt", 1234)
	// err = c.AddInboundClient(2, inboundClient)
	// if err != nil {
	// 	logger.Error("Creating client failed", slog.String("error", err.Error()))
	// 	return
	// }

	// // Generate default inbound configuration
	// defaultInbound, err := c.GenerateDefaultInboundConfig("testinbound", "tori.fi", "178.236.244.241", 444)
	// if err != nil {
	// 	logger.Error("Failed to generate default inbound config", "error", err)
	// 	return
	// }

	// // Add the inbound
	// _, err = c.AddInbound(defaultInbound)
	// if err != nil {
	// 	logger.Error("Failed to add inbound", "error", err)
	// 	return
	// }

}

func newLogger(level string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	default:
		logLevel = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	return slog.New(handler)
}
