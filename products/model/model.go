package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        *string            `json:"name" bson:"name"`
	Pictures    *[]string          `json:"pictures" bson:"pictures"`
	PromotionID *int64             `json:"promotion_id" bson:"promotion_id"`
	Price       *int64             `json:"price" bson:"price"`
	Quantity    *uint64            `json:"quantity" bson:"quantity"`
	Unit        string             `json:"unit" bson:"unit"`
	Description *string            `json:"description" bson:"description"`
	Detail      *string            `json:"detail" bson:"detail"`
	Categories  *[]string          `json:"categories" bson:"categories"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type CreateProductRequest struct {
	Name        *string   `json:"name" bson:"name"`
	Pictures    *[]string `json:"pictures" bson:"pictures"`
	Price       *int64    `json:"price" bson:"price"`
	Quantity    *uint64   `json:"quantity" bson:"quantity"`
	Description *string   `json:"description" bson:"description"`
	Detail      *string   `json:"detail" bson:"detail"`
	Categories  *[]string `json:"categories" bson:"categories"`
	Unit        *string   `json:"unit" bson:"unit"`
}

const (
	ColCategory = "categories"
)
