package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"videoai/internal/middleware"
	"github.com/gin-gonic/gin"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedType   string
	}{
		{
			name:           "Successful health check",
			queryParam:     "",
			expectedStatus: http.StatusOK,
			expectedType:   "success",
		},
		{
			name:           "Error health check",
			queryParam:     "?error=true",
			expectedStatus: http.StatusInternalServerError,
			expectedType:   "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(middleware.ErrorHandler())
			r.GET("/health", HealthCheck)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/health"+tt.queryParam, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to parse JSON response: %v", err)
			}

			if response["status"] != tt.expectedType {
				t.Errorf("Expected status type %s, got %s", tt.expectedType, response["status"])
			}
		})
	}	
}