package repository

import (
	"context"

	"github.com/tristaamne/flowershopbe-v4/products/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	GetProductByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.Product, error)
	CreateAProduct(ctx context.Context, data *model.Product) (primitive.ObjectID, error)
	DeleteAProduct(ctx context.Context, filter bson.M, opts *options.DeleteOptions) error
	UpdateAProduct(ctx context.Context, filter bson.M, data bson.M) error
}
