package responses

import (
	"net/http"

	"videoai/internal/errors"

	"github.com/gin-gonic/gin"
)

// Success sends a JSend success response
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Error sends a JSend error or fail response
func Error(c *gin.Context, err *errors.APIError) {
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"code":   http.StatusInternalServerError,
			"message": errors.MsgInternalError,
		})
		return
	}
	
	response := gin.H{
		"status": err.Type,
		"code":   err.Code,
	}
	
	if err.ClientMessage != "" {
		response["message"] = err.ClientMessage
	}
	
	if err.Data != nil {
		response["data"] = err.Data
	}
	
	c.JSON(err.Code, response)
}
