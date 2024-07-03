package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
}

func Init() Config {
	return Config{
		Port:     os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
