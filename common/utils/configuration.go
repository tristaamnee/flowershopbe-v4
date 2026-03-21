package utils

import (
	"os"

	"github.com/tristaamne/flowershopbe-v4/common/db"
)

func LoadConfig() db.DatabaseConfiguration {
	return db.DatabaseConfiguration{
		//User:     os.Getenv("DB_USER"),
		//Password: os.Getenv("DB_PASSWORD"),
		Address:  os.Getenv("DB_ADDRESS"),
		Database: os.Getenv("DB_NAME"),
	}
}
