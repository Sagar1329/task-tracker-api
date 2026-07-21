package routes

import (
	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, app *app.Application) {

	healthHandler := handlers.NewHealthHandler(app)

	router.GET("/health", healthHandler.Health)

	router.GET("/db-health", healthHandler.DBHealth)
}