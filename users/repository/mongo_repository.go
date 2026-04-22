package repository

import (
	"context"

	base "github.com/tristaamne/flowershopbe-v4/common/repository/mongodb"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) UserRepository {
	return &userRepository{coll: coll}
}

func (r *userRepository) GetUserByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.User, error) {
	return base.GetByCondition[model.User](ctx, r.coll, filter, opts)
}

func (r *userRepository) RegisterUser(ctx context.Context, user *model.User) (primitive.ObjectID, error) {
	return base.Create(ctx, r.coll, user)
}

func (r *userRepository) DeleteAUser(ctx context.Context, filter bson.M, deleteOptions *options.DeleteOptions) error {
	return base.DeleteByCondition(ctx, r.coll, filter, deleteOptions)
}

func (r *userRepository) UpdateAUser(ctx context.Context, filter bson.M, update bson.M) error {
	return base.UpdateByCondition(ctx, r.coll, filter, update)
}
