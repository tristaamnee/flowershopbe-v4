package repository

import (
	base "github.com/tristaamne/flowershopbe-v4/common/repository"
	"github.com/tristaamne/flowershopbe-v4/token/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateToken(coll *mongo.Collection, token *model.Token) (primitive.ObjectID, error) {
	return base.Create(coll, token)
}

func GetToken(coll *mongo.Collection, filter bson.M, opts *options.FindOptions) ([]model.Token, error) {
	return base.GetByCondition[model.Token](coll, filter, opts)
}

func DeleteToken(coll *mongo.Collection, filter bson.M, opts *options.DeleteOptions) error {
	return base.DeleteByCondition(coll, filter, opts)
}
