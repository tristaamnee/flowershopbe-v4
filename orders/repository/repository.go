package repository

import (
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByCondition(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]model.Order, error) {
	return base.GetByCondition[model.Order](coll, filter, opts)
}

func CreateAOrder(coll *mongo.Collection, data *model.Order) (primitive.ObjectID, error) {
	return base.Create(coll, data)
}

func DeleteAOrder(coll *mongo.Collection, filter bson.M, opts *options.DeleteOptions) error {
	return base.DeleteByCondition(coll, filter, opts)
}

func UpdateAOrder(coll *mongo.Collection, filter bson.M, data bson.M) error {
	return base.UpdateByCondition(coll, filter, data)
}
