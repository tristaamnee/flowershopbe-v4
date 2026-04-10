package model

import (
	userModel "github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID              primitive.ObjectID        `json:"id" bson:"_id,omitempty"`
	UserID          primitive.ObjectID        `json:"user_id" bson:"user_id,omitempty"`
	Items           []OrderItem               `json:"items" bson:"items,omitempty"`
	PromotionIDs    []string                  `json:"promotion_ids" bson:"promotion_ids,omitempty"`
	TotalPrice      int64                     `json:"total_price" bson:"total_price,omitempty"`
	DeliveryAddress userModel.DeliveryAddress `json:"delivery_address" bson:"delivery_address,omitempty"`
	// Status: pending, processing, shipping, completed, canceled
	Status    string    `json:"status" bson:"status,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	Quantity  int                `json:"quantity" bson:"quantity,omitempty"`
	Price     int64              `json:"price" bson:"price,omitempty"`
}
type OrderRequest struct {
	Items           []OrderItem               `json:"items" bson:"items,omitempty"`
	PromotionIDs    []primitive.ObjectID      `json:"promotionIds" bson:"promotionIds,omitempty"`
	DeliveryAddress userModel.DeliveryAddress `json:"delivery_address" bson:"delivery_address,omitempty"`
}
