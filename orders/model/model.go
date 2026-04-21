package model

import (
	"time"

	userModel "github.com/tristaamne/flowershopbe-v4/users/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	Quantity  int                `json:"quantity" bson:"quantity,omitempty"`
	Price     int64              `json:"price" bson:"price,omitempty"`
}

type Order struct {
	ID              primitive.ObjectID        `json:"id" bson:"_id,omitempty"`
	OrderNumber     int64                     `json:"order_number" bson:"order_number"`
	UserName        string                    `json:"user_name" bson:"user_name"`
	OrderItems      []OrderItem               `json:"order_items" bson:"order_items"`
	PromotionIDs    []primitive.ObjectID      `json:"promotion_ids" bson:"promotion_ids"`
	TotalPrice      int64                     `json:"total_price" bson:"total_price"`
	DeliveryAddress userModel.DeliveryAddress `json:"delivery_address" bson:"delivery_address"`
	// Status: pending, processing, shipping, completed, canceled
	Status           string    `json:"status" bson:"status,omitempty"`
	PaymentSignature string    `json:"payment_signature" bson:"payment_signature"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

type OrderRequest struct {
	OrderItems      []OrderItem               `json:"order_items" bson:"order_items,omitempty"`
	PromotionIDs    []primitive.ObjectID      `json:"promotion_ids" bson:"promotion_ids"`
	DeliveryAddress userModel.DeliveryAddress `json:"delivery_address" bson:"delivery_address"`
}

const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusShipping   = "shipping"
	StatusComplete   = "complete"
	StatusCancelled  = "cancelled"
)
