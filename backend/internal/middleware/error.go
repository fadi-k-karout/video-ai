package middleware

import (
	"log/slog"

	"videoai/internal/errors"
	"videoai/internal/responses"

	"github.com/gin-gonic/gin"
)

// ErrorHandler processes errors and returns JSend-compliant responses
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			apiErr := errors.GetAPIError(err)

			// Log error with request context
			requestID, ok := c.Get("request_id")
			if !ok {
				requestID = "unknown"
			}

			slog.Error("request_error",
				"error", err.Error(),
				"status_code", apiErr.Code,
				"method", c.Request.Method,
				SafeAttr("path", c.Request.URL.Path),
				"request_id", requestID,
				"client_ip", c.ClientIP(),
			)

			responses.Error(c, apiErr)
			c.Abort()
		}
	}
}
