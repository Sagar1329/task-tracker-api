package app

import (
	"github.com/Sagar1329/task-tracker-api/internal/config"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DB     *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) *Application {
	return &Application{
		Config: cfg,
		DB:     db,
	}
}