package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"videoai/internal/errors"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		setupError     func(*gin.Context)
		expectedStatus int
		expectedType   string
		expectedMsg    string
	}{
		{
			name: "BadRequest error",
			setupError: func(c *gin.Context) {
				c.Error(errors.BadRequest("invalid input", nil))
			},
			expectedStatus: http.StatusBadRequest,
			expectedType:   "fail",
			expectedMsg:    errors.MsgBadRequest,
		},
		{
			name: "Unauthorized error",
			setupError: func(c *gin.Context) {
				c.Error(errors.Unauthorized("access denied", nil))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedType:   "fail",
			expectedMsg:    errors.MsgUnauthorized,
		},
		{
			name: "NotFound error",
			setupError: func(c *gin.Context) {
				c.Error(errors.NotFound("resource not found", nil))
			},
			expectedStatus: http.StatusNotFound,
			expectedType:   "fail",
			expectedMsg:    errors.MsgNotFound,
		},
		{
			name: "InternalError",
			setupError: func(c *gin.Context) {
				c.Error(errors.InternalError("server error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedType:   "error",
			expectedMsg:    errors.MsgInternalError,
		},
		{
			name: "No error - should not interfere",
			setupError: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "success"})
			},
			expectedStatus: http.StatusOK,
			expectedType:   "",
			expectedMsg:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := gin.New()
			r.Use(ErrorHandler())
			
			r.GET("/test", func(c *gin.Context) {
				c.Set("request_id", "test-request-id")
				tt.setupError(c)
			})

			req := httptest.NewRequest("GET", "/test", nil)
			r.ServeHTTP(w, req)

			// Verify response
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedType != "" {
				var response map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response["status"] != tt.expectedType {
					t.Errorf("Expected status type %s, got %s", tt.expectedType, response["status"])
				}

				if response["message"] != tt.expectedMsg {
					t.Errorf("Expected message %s, got %s", tt.expectedMsg, response["message"])
				}
			}
		})
	}
}
