package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Success(
	c *gin.Context,
	statusCode int,
	message string,
	data any,
) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(
	c *gin.Context,
	statusCode int,
	message string,
) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}