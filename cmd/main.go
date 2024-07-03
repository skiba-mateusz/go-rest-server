package main

import (
	"fmt"
	"log"

	"github.com/skiba-mateusz/go-rest-server/cmd/api"
	"github.com/skiba-mateusz/go-rest-server/config"
	"github.com/skiba-mateusz/go-rest-server/database"
)

func main() {
	cfg := config.Init()

	_, err := database.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Fatal("Could not connect DB: ", err)
	}
	log.Println("DB connected")

	server := api.NewAPIServer(fmt.Sprintf(":%s", cfg.Port))
	if err := server.Run(); err != nil {
		log.Fatal("Could not run server: ", err)
	}
}
