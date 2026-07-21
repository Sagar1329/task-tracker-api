package main

import (
	"fmt"
	"log"
			"github.com/Sagar1329/task-tracker-api/internal/app"

	"github.com/Sagar1329/task-tracker-api/internal/database"
	"github.com/Sagar1329/task-tracker-api/internal/config"
		"github.com/Sagar1329/task-tracker-api/internal/server"

)

func main(){
	cfg := config.Load()

	fmt.Println("Application:", cfg.AppName)
	fmt.Println("Running on Port:", cfg.AppPort)
	fmt.Println("Database", cfg.DBName)
	fmt.Println("Host:", cfg.DBHost)
	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Application Started")
	application := app.New(cfg, db)
	server := server.New(application)
	server.Start()

}