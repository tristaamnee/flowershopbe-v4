package repository

import (
	"context"
	"log"
	"time"

	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByCondition(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		er := cursor.Close(ctx)
		if er != nil {
			log.Printf("Error while closing the cursor: %v", er)
		}
	}(cursor, ctx)

	var data []bson.M

	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	result, err := utils.ConvertBsonToJson(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
