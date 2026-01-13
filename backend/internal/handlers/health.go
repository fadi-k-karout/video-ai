package handlers

import (
	"github.com/gin-gonic/gin"
	"videoai/internal/errors"
	"videoai/internal/responses"
)

func HealthCheck(c *gin.Context) {
	// Simulate error condition for demonstration
	if c.Query("error") == "true" {
		c.Error(errors.InternalError("health check failed"))
		return
	}

	responses.Success(c, gin.H{
		"message": "Video AI server is running",
		"version": "1.0.0",
	})
}