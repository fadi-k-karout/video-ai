package main

import (
	"log"

	"videoai/internal/config"
	"videoai/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	router := routes.SetupRoutes(cfg)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run("0.0.0.0:" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
