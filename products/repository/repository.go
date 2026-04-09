package repository

import (
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProductByCondition(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]model.Product, error) {
	return base.GetByCondition[model.Product](coll, filter, opts)
}

func CreateAProduct(coll *mongo.Collection, data *model.Product) (primitive.ObjectID, error) {
	return base.Create(coll, data)
}

func DeleteAProduct(coll *mongo.Collection, filter bson.M, opts *options.DeleteOptions) error {
	return base.DeleteByCondition(coll, filter, opts)
}

func UpdateAProduct(coll *mongo.Collection, filter bson.M, data bson.M) error {
	return base.UpdateByCondition(coll, filter, data)
}
