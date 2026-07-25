package routes

import (
	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/handlers"
		"github.com/Sagar1329/task-tracker-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, app *app.Application) {

	healthHandler := handlers.NewHealthHandler(app)
	authHandler := handlers.NewAuthHandler(app)
jobApplicationHandler := handlers.NewJobApplicationHandler(app)
	api := router.Group("/api/v1")

	health := api.Group("/health")
	{
	health.GET("/", healthHandler.Health)

	health.GET("/db", healthHandler.DBHealth)
	}

	auth := api.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}
    jobApplications := api.Group("/job-applications")
     jobApplications.Use(middleware.AuthMiddleware(app.Config.JWT.Secret))
{
	jobApplications.POST("/", jobApplicationHandler.Create)
	jobApplications.GET("/", jobApplicationHandler.GetAll)
	jobApplications.GET("/:id", jobApplicationHandler.GetByID)
	jobApplications.PUT("/:id", jobApplicationHandler.Update)
	jobApplications.DELETE("/:id", jobApplicationHandler.Delete)
}

}