package db

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DatabaseConfiguration struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Database string `json:"database"`
}

func ConnectDatabase(config DatabaseConfiguration) *pg.DB {
	opts := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Address,
		Database: config.Database,
	}
	return pg.Connect(opts)
}

func CreateTable(db *pg.DB, model interface{}) error {
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
		Temp:        false,
	})
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase(db *pg.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
