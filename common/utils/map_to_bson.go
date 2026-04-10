package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapToBSon(data map[string]interface{}) bson.M {
	filter := bson.M{}
	for k, v := range data {
		if id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", v)); err == nil {
			filter[k] = id
		} else {
			filter[k] = v
		}
	}
	return filter
}
