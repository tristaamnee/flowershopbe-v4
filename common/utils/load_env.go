package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	er := godotenv.Load()
	if er != nil {
		log.Println("Warning: No .env file found")
	}
}
