package repository

import (
	"context"

	base "github.com/tristaamne/flowershopbe-v4/common/repository/mongodb"
	"github.com/tristaamne/flowershopbe-v4/orders/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type orderRepo struct {
	coll *mongo.Collection
}

func NewMongoOrderRepository(coll *mongo.Collection) OrderRepository {
	return &orderRepo{coll: coll}
}

func (r *orderRepo) GetByCondition(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]model.Order, error) {
	return base.GetByCondition[model.Order](ctx, r.coll, filter, opts)
}

func (r *orderRepo) CreateAOrder(ctx context.Context, data *model.Order) (primitive.ObjectID, error) {
	return base.Create(ctx, r.coll, data)
}

func (r *orderRepo) DeleteAOrder(ctx context.Context, filter bson.M, opts *options.DeleteOptions) error {
	return base.DeleteByCondition(ctx, r.coll, filter, opts)
}

func (r *orderRepo) UpdateAOrder(ctx context.Context, filter bson.M, data bson.M) error {
	return base.UpdateByCondition(ctx, r.coll, filter, data)
}
