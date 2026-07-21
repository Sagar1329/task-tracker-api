package server

import (
	"log"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app    *app.Application
	router *gin.Engine
}

func New(app *app.Application) *Server {

	router := gin.Default()

	routes.Register(router, app)

	return &Server{
		app:    app,
		router: router,
	}
}

func (s *Server) Start() {

	log.Printf("Server running on port %s", s.app.Config.AppPort)

	err := s.router.Run(":" + s.app.Config.AppPort)

	if err != nil {
		log.Fatal(err)
	}
}