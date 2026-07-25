package main

import (
	"os"
	"log/slog"

	"github.com/Sagar1329/task-tracker-api/internal/app"
	"github.com/Sagar1329/task-tracker-api/internal/config"
	"github.com/Sagar1329/task-tracker-api/internal/database"
	"github.com/Sagar1329/task-tracker-api/internal/logger"
	"github.com/Sagar1329/task-tracker-api/internal/server"
)

func main() {

	// Initialize Logger
	logger.Init()

	// Load Configuration
	cfg := config.Load()

	// Validate Configuration
	if err := cfg.Validate(); err != nil {
		logger.Log.Error(
	"Application startup failed",
	slog.Any("error", err),
)
	}

	logger.Log.Info(
		"Configuration loaded successfully",
		slog.String("application", cfg.AppName),
		slog.String("port", cfg.Server.Port),
		slog.String("database", cfg.Database.Name),
		slog.String("host", cfg.Database.Host),
	)

	// Connect Database
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Log.Error(
			"Failed to connect to database",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	logger.Log.Info("Database connected successfully")

	// Create Application
	application := app.New(cfg, db)

	// Create Server
	srv := server.New(application)

	logger.Log.Info(
		"Starting HTTP server",
		slog.String("port", cfg.Server.Port),
	)

	// Start Server
	srv.Start()
}