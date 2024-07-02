package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skiba-mateusz/go-rest-server/cmd/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")

	server := api.NewAPIServer(listenAddr)
	if err := server.Run(); err != nil {
		log.Fatal("Could not run server: ", err)
	}
}
