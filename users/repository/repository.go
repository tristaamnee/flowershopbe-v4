package repository

import (
	"context"

	"github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetUserByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.User, error)
	RegisterUser(ctx context.Context, user *model.User) (primitive.ObjectID, error)
	DeleteAUser(ctx context.Context, filter bson.M, deleteOptions *options.DeleteOptions) error
	UpdateAUser(ctx context.Context, filter bson.M, update bson.M) error
}
