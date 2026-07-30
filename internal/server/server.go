package server

import (
	"time"
	"github.com/Sagar1329/task-tracker-api/internal/logger"
"log/slog"
"os"
	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/routes"
	"github.com/gin-gonic/gin"
	 "github.com/gin-contrib/cors"
)

type Server struct {
	app    *app.Application
	router *gin.Engine
}

func New(app *app.Application) *Server {

	router := gin.Default()
router.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:3000",
				"http://localhost:3001",

	},

	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	},

	AllowHeaders: []string{
		"Origin",
		"Content-Type",
		"Authorization",
	},

	ExposeHeaders: []string{
		"Content-Length",
	},

	AllowCredentials: true,

	MaxAge: 12 * time.Hour,
}))
	routes.Register(router, app) 

	return &Server{
		app:    app,
		router: router,
	}
}

func (s *Server) Start() {

	logger.Log.Info("Server running on port %s", s.app.Config.Server.Port)

	err := s.router.Run(":" + s.app.Config.Server.Port)

	if err != nil {
		logger.Log.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)	 
	}
}