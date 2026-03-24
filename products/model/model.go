package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Pictures    []string           `json:"pictures" bson:"pictures"`
	PromotionID int64              `json:"promotion_id" bson:"promotion_id"`
	Price       int64              `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`
	Detail      string             `json:"detail" bson:"detail"`
	Categories  []string           `json:"categories" bson:"categories"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	_           struct{}           `pg:"_schema:products"`
}

type CreateProductRequest struct {
	Name        string   `json:"name" bson:"name"`
	Pictures    []string `json:"pictures" bson:"pictures"`
	Price       int64    `json:"price" bson:"price"`
	Description string   `json:"description" bson:"description"`
	Detail      string   `json:"detail" bson:"detail"`
	Categories  []string `json:"categories" bson:"categories"`
}

type UpdateProductRequest struct {
	Name        *string   `json:"name"`
	Pictures    *[]string `json:"pictures"`
	Description *string   `json:"description"`
	Detail      *string   `json:"detail"`
	Categories  *[]string `json:"categories"`
	Price       *int64    `json:"price"`
}

const (
	ColCategory = "categories"
)
