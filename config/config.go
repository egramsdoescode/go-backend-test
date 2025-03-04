package config

import (
	"log"

	env "github.com/joho/godotenv"
)

func LoadEnv() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables")
	}
}
