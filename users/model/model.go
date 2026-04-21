package model

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name              string             `json:"name" bson:"name"`
	Password          string             `json:"-" bson:"password"`
	Birthday          time.Time          `json:"birthday" bson:"birthday"`
	Email             string             `json:"email" bson:"email"`
	DeliveryAddresses []DeliveryAddress  `json:"delivery_addresses" bson:"delivery_addresses"`
	ProfilePicture    string             `json:"profile_picture" bson:"profile_picture"`
	Role              int                `json:"role" bson:"role"`
	ProviderID        string             `bson:"provider_id"`
	EmailVerified     bool               `json:"email_verified" bson:"email_verified"`
	UpdatedAt         time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
}

type UserRequest struct {
	Name              string                `json:"name" form:"name"`
	Password          string                `json:"password" form:"password"`
	Birthday          time.Time             `json:"birthday" form:"birthday" time_format:"2006-01-02"`
	Email             string                `json:"email" form:"email"`
	DeliveryAddresses []DeliveryAddress     `json:"delivery_addresses" form:"delivery_addresses"`
	ProfilePicture    *multipart.FileHeader `form:"profile_picture"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
}

type DeliveryAddress struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ReceiverName string             `json:"receiver_name" bson:"receiver_name"`
	Address      string             `json:"address" bson:"address"`
	Phone        string             `json:"phone" bson:"phone"`
	IsDefault    bool               `json:"is_default" bson:"is_default"`
}

const (
	ColEmail = "email"
)
