package repository

import (
	"context"

	base "github.com/tristaamne/flowershopbe-v4/common/repository/mongodb"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type prodRepository struct {
	coll *mongo.Collection
}

func NewProductRepository(coll *mongo.Collection) ProductRepository {
	return &prodRepository{coll: coll}
}

func (r *prodRepository) GetProductByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.Product, error) {
	return base.GetByCondition[model.Product](ctx, r.coll, filter, opts)
}

func (r *prodRepository) CreateAProduct(ctx context.Context, data *model.Product) (primitive.ObjectID, error) {
	return base.Create(ctx, r.coll, data)
}

func (r *prodRepository) DeleteAProduct(ctx context.Context, filter bson.M, opts *options.DeleteOptions) error {
	return base.DeleteByCondition(ctx, r.coll, filter, opts)
}

func (r *prodRepository) UpdateAProduct(ctx context.Context, filter bson.M, data bson.M) error {
	return base.UpdateByCondition(ctx, r.coll, filter, data)
}
