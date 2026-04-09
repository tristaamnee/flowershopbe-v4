package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID          primitive.ObjectID   `json:"userId" bson:"userId,omitempty"`
	FlowerIDs       []primitive.ObjectID `json:"flowerIds" bson:"flowerIds,omitempty"`
	PromotionIDs    []primitive.ObjectID `json:"promotionIds" bson:"promotionIds,omitempty"`
	TotalPrice      int64                `json:"totalPrice" bson:"totalPrice,omitempty"`
	DeliveryAddress string               `json:"deliveryAddress" bson:"deliveryAddress,omitempty"`
	Status          string               `json:"status" bson:"status,omitempty"`
	UpdatedAt       time.Time            `json:"updatedAt" bson:"updatedAt,omitempty"`
	CreatedAt       time.Time            `json:"createdAt" bson:"createdAt,omitempty"`
}

type OrderRequest struct {
	FlowerIDs       []primitive.ObjectID `json:"flowerIds" bson:"flowerIds,omitempty"`
	PromotionIDs    []primitive.ObjectID `json:"promotionIds" bson:"promotionIds,omitempty"`
	TotalPrice      int64                `json:"totalPrice" bson:"totalPrice,omitempty"`
	DeliveryAddress string               `json:"deliveryAddress" bson:"deliveryAddress,omitempty"`
}
