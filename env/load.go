package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	ENVIRONMENT string
}

var DefaultConfig Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("$PORT must be set")
	}

	env := os.Getenv("ENVIRONMENT")
	if port == "" {
		log.Fatalln("$ENVIRONMENT must be set")
	}

	DefaultConfig = Config{
		PORT:        port,
		ENVIRONMENT: env,
	}

	log.Println("Loaded .env!")
}
