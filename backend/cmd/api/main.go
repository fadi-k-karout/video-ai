package main

import (
	"log/slog"
	"os"

	"videoai/internal/config"
	stdlog "videoai/internal/logger"
	"videoai/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration: ", "error", err)
		os.Exit(1)
	}
    // we are setting the logger and the default slog handler
	// we are returning an instance if later needed for DI
	_ = stdlog.Init(cfg)

	slog.Info("Starting Video AI server",
		"port", cfg.Port,
		"environment", cfg.Environment,
		"log_level", cfg.LogLevel,
		"log_format", cfg.LogFormat,
	)

	router := routes.SetupRoutes(cfg)

	if err := router.Run("0.0.0.0:" + cfg.Port); err != nil {
		slog.Error("Failed to start server: ", "error", err)
		os.Exit(1)
	}
}
