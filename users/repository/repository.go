package repository

import (
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserByCondition(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]model.User, error) {
	return base.GetByCondition[model.User](coll, filter, opts)
}

func RegisterUser(coll *mongo.Collection, user *model.User) (primitive.ObjectID, error) {
	return base.Create(coll, user)
}

func DeleteAUser(coll *mongo.Collection, filter bson.M, deleteOptions *options.DeleteOptions) error {
	return base.DeleteByCondition(coll, filter, deleteOptions)
}

func UpdateAUser(coll *mongo.Collection, filter bson.M, update bson.M) error {
	return base.UpdateByCondition(coll, filter, update)
}
