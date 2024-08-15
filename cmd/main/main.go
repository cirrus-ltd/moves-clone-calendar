package main

import (
	"log"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/config"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/db"
	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/infrastructure/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: ")
	}
	router.InitRouter()
}
