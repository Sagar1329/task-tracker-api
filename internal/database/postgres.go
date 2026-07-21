package database

import (
	"fmt"
	"log"

	"github.com/Sagar1329/task-tracker-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connect(cfg *config.Config)(*gorm.DB, error){
	dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
sqlDB, err := db.DB()
if err != nil {
	return nil, err
}

if err := sqlDB.Ping(); err != nil {
	return nil, err
}
	log.Println("Connected to PostgreSQL")

	return db, nil

}