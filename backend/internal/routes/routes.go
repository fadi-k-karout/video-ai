package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"videoai/internal/config"
	"videoai/internal/handlers"
	"videoai/internal/middleware"
)

func SetupRoutes(cfg *config.Config) *gin.Engine {
	if cfg.Environment == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// Logging middleware: structured in prod, default gin logger in dev
	if cfg.Environment == config.EnvProduction {
		r.Use(middleware.RequestLogger())
	} else {
		r.Use(gin.Logger())
	}

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Error handling middleware
	r.Use(middleware.ErrorHandler())

	// API routes
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
	}

	return r
}
