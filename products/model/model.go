package model

import "time"

type Product struct {
	ID          int64     `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Pictures    []string  `json:"pictures" bson:"pictures"`
	PromotionID int64     `json:"promotion_id" bson:"promotion_id"`
	Price       int64     `json:"price" bson:"price"`
	Description string    `json:"description" bson:"description"`
	Detail      string    `json:"detail" bson:"detail"`
	Categories  []string  `json:"categories" bson:"categories"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	_           struct{}  `pg:"_schema:products"`
}

const (
	ColCategory = "categories"
)
