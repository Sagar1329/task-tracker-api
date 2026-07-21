package handlers

import (
	"net/http"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	app *app.Application
}

func NewHealthHandler(app *app.Application) *HealthHandler {
	return &HealthHandler{
		app: app,
	}
}
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
func (h *HealthHandler) DBHealth(c *gin.Context) {
sqlDB, err := h.app.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Database Error",
			"error": err.Error(),
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Database Unreachable",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Database Connected",
	})
}