package utils

import (
	"encoding/json"
	"os"

	"github.com/tristaamne/flowershopbe-v4/common/db"
)

func Get() (interface{}, error) {
	file, err := os.Open("src/configturation.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config db.DatabaseConfiguration

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
