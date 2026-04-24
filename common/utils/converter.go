package utils

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertBsonToJson(b []bson.M) (json.RawMessage, error) {
	j, err := bson.MarshalExtJSON(b, false, false)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func ConvertStringToID(idStr string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objID, nil
}
