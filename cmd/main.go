package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skiba-mateusz/go-rest-server/cmd/api"
	"github.com/skiba-mateusz/go-rest-server/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	_, err := database.NewMongoClient(mongoURI)
	if err != nil {
		log.Fatal("Could not connect DB: ", err)
	}
	log.Println("DB connected")

	server := api.NewAPIServer(fmt.Sprintf(":%s", port))
	if err := server.Run(); err != nil {
		log.Fatal("Could not run server: ", err)
	}
}
