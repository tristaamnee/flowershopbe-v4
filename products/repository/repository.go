package repository

import (
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByCategory(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) (interface{}, error) {
	return base.GetByCondition(coll, filter, opts)
}

func CreateAProduct(coll *mongo.Collection, data interface{}) (interface{}, error) {
	return base.Create(coll, data)
}
