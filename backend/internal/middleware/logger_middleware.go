package middleware

import (
	"log/slog"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SafeAttr creates a sanitized slog attribute for user-controlled values
func SafeAttr(key, value string) slog.Attr {
	sanitized := strings.Map(func(r rune) rune {
		// Remove control characters except tab (which is often valid)
		if unicode.IsControl(r) && r != '\t' {
			return -1
		}
		return r
	}, value)
	
	// Truncate excessively long values
    if len([]rune(sanitized)) > 500 {
		sanitized = string([]rune(sanitized)[:500]) + "..."
 	}	
	
	return slog.String(key, sanitized)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		requestID := uuid.New().String()
		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)

		c.Next()

		duration := time.Since(start)
		
		slog.Info("request",
			"method", c.Request.Method,
			SafeAttr("path", c.Request.URL.Path),
			"status", c.Writer.Status(),
			"duration_ms", duration.Milliseconds(),
			"client_ip", c.ClientIP(),
			SafeAttr("user_agent", c.Request.UserAgent()),
			"request_id", requestID,
		)
	}
}