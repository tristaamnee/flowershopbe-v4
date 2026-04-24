package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/products/model"
	"github.com/tristaamne/flowershopbe-v4/products/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	repo repository.ProductRepository
	cfg  *config.Config
}

type Service interface {
	CreateANewProduct(ctx context.Context, req model.CreateProductRequest) (primitive.ObjectID, error)
	DeleteProductByID(ctx context.Context, id primitive.ObjectID) error
	UpdateAProduct(ctx context.Context, req model.CreateProductRequest, id primitive.ObjectID) (primitive.ObjectID, error)
	UpdateAProductQuantity(ctx context.Context, rollbackData map[primitive.ObjectID]interface{}, id primitive.ObjectID, amount *uint64, changeType bool) (primitive.ObjectID, error)
	GetProductByCategory(ctx context.Context, category, pageStr, limitStr, sortField, orderStr string) ([]model.Product, error)
	GetProductByID(ctx context.Context, ids []primitive.ObjectID) ([]model.Product, error)
}

func NewService(repo repository.ProductRepository, cfg *config.Config) Service {
	return &service{repo: repo, cfg: cfg}
}

func (s *service) CreateANewProduct(ctx context.Context, req model.CreateProductRequest) (primitive.ObjectID, error) {
	product := &model.Product{
		Name:        req.Name,
		Pictures:    req.Pictures,
		Description: req.Description,
		Price:       req.Price,
		Detail:      req.Detail,
		Categories:  req.Categories,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return s.repo.CreateAProduct(ctx, product)
}

func (s *service) DeleteProductByID(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	opts := options.Delete()

	err := s.repo.DeleteAProduct(ctx, filter, opts)
	if err != nil {
		return fmt.Errorf("DeleteProductByID: %w", err)
	}
	return nil
}

func (s *service) UpdateAProduct(ctx context.Context, req model.CreateProductRequest, id primitive.ObjectID) (primitive.ObjectID, error) {
	update := bson.M{}

	if req.Name != nil {
		update["name"] = *req.Name
	}
	if req.Price != nil {
		update["price"] = req.Price
	}
	if req.Description != nil {
		update["description"] = *req.Description
	}
	if req.Detail != nil {
		update["detail"] = *req.Detail
	}
	if req.Categories != nil {
		update["categories"] = *req.Categories
	}
	if req.Unit != nil {
		update["unit"] = *req.Unit
	}

	if len(update) == 0 {
		return id, fmt.Errorf("no update found to update")
	}

	update["updated_at"] = time.Now()

	update = bson.M{
		"$set": update,
	}

	filter := bson.M{"_id": id}
	er := s.repo.UpdateAProduct(ctx, filter, update)
	if er != nil {
		return primitive.NilObjectID, fmt.Errorf("error when update a product: %w", er)
	}
	return id, nil
}

func (s *service) UpdateAProductQuantity(ctx context.Context, rollbackData map[primitive.ObjectID]interface{}, id primitive.ObjectID, amountAddr *uint64, changeType bool) (primitive.ObjectID, error) {
	// type true = increase
	// type false = decrease
	filter := bson.M{"_id": id}
	ids := []primitive.ObjectID{id}

	product, err := s.GetProductByID(ctx, ids)
	if err != nil || len(product) == 0 {
		return id, err
	}

	quantity := *product[0].Quantity
	amount := *amountAddr

	if product[0].Quantity == nil || (quantity < amount && !changeType) {
		return id, fmt.Errorf("not enough product quantity")
	}

	var changeAmount int64
	if changeType {
		changeAmount = int64(amount)
	} else {
		changeAmount = -int64(amount)
		filter["quantity"] = bson.M{"$gte": amount}
	}

	update := bson.M{
		"$inc": bson.M{"quantity": changeAmount},
		"$set": bson.M{"updated_at": time.Now()},
	}

	err = s.repo.UpdateAProduct(ctx, filter, update)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("error when update a product: %w", err)
	}
	rollbackData[id] = amount

	return id, nil
}

func (s *service) GetProductByCategory(ctx context.Context, category, pageStr, limitStr, sortField, orderStr string) ([]model.Product, error) {
	filter := bson.M{}
	opts := options.Find()

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	order, err := strconv.Atoi(orderStr)
	if err != nil || (order != 1 && order != -1) {
		order = -1
	}

	allowedSortFields := map[string]bool{
		"create_at": true,
		"price":     true,
		"name":      true,
	}

	if !allowedSortFields[sortField] {
		sortField = "create_at"
	}

	skip := (page - 1) * limit

	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))
	opts.SetSort(bson.D{{sortField, order}})

	if category != "" {
		filter[model.ColCategory] = bson.M{
			"$elemMatch": bson.M{
				"$regex":   category,
				"$options": "i"},
		}
	}

	data, err := s.repo.GetProductByCondition(ctx, filter, opts)
	if err != nil {
		return []model.Product{}, fmt.Errorf("GetProductByCondition: %w", err)
	}
	return data, nil
}

func (s *service) GetProductByID(ctx context.Context, ids []primitive.ObjectID) ([]model.Product, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	opts := options.Find()

	productData, err := s.repo.GetProductByCondition(ctx, filter, opts)
	if err != nil {
		return []model.Product{}, fmt.Errorf("GetProductByID: %w", err)
	}
	if len(productData) == 0 {
		return []model.Product{}, nil
	}
	return productData, nil
}
