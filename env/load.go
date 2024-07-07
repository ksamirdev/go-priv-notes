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
		log.Fatalf("[env] Error loading file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("[env] $PORT must be set")
	}

	env := os.Getenv("ENVIRONMENT")
	if port == "" {
		log.Fatalln("[env] $ENVIRONMENT must be set")
	}

	DefaultConfig = Config{
		PORT:        port,
		ENVIRONMENT: env,
	}

	log.Println("[env] loaded!")
}
