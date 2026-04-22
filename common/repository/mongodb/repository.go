package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByCondition[T any](ctx context.Context, coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]T, error) {

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func Create[T any](ctx context.Context, coll *mongo.Collection, data T) (primitive.ObjectID, error) {

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

func DeleteByCondition(ctx context.Context, coll *mongo.Collection, filter bson.M, opts *options.DeleteOptions) error {

	_, err := coll.DeleteOne(ctx, filter, opts)
	if err != nil {
		return err
	}
	return nil
}

func UpdateByCondition[T any](ctx context.Context, coll *mongo.Collection, filter bson.M, data T) error {

	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}
