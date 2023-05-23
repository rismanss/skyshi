package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
}
