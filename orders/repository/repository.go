package repository

import (
	"context"

	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository interface {
	CreateAOrder(ctx context.Context, data *model.Order) (primitive.ObjectID, error)
	GetByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.Order, error)
	DeleteAOrder(ctx context.Context, filter bson.M, opts *options.DeleteOptions) error
	UpdateAOrder(ctx context.Context, filter bson.M, data bson.M) error
}
