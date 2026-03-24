package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByCondition[T any](coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error when close cursor in GetByCondition: %v", err)
			return
		}
	}(cursor, ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func Create[T any](coll *mongo.Collection, data T) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("fail to convert InsertedID to primitive.ObjectID")
	}
	return id, nil
}

func DeleteByCondition(coll *mongo.Collection, filter bson.M, opts *options.DeleteOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := coll.DeleteOne(ctx, filter, opts)
	if err != nil {
		return err
	}
	return nil
}

func UpdateByCondition[T any](coll *mongo.Collection, filter bson.M, data T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}
