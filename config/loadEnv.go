package config

import (
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
