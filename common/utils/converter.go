package utils

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertBsonToJson(b []bson.M) (json.RawMessage, error) {
	j, err := bson.MarshalExtJSON(b, false, false)
	if err != nil {
		return nil, err
	}
	return j, nil
}
