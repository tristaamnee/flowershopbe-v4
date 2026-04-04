package model

import (
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
	PhoneNumber       string             `json:"phone_number" bson:"phone_number"`
	ProfilePicture    string             `json:"profile_picture" bson:"profile_picture"`
	Role              int64              `json:"role" bson:"role"`
	UpdatedAt         time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
}

type UserRequest struct {
	Name              string            `json:"name" bson:"name"`
	Password          string            `json:"-" bson:"password"`
	Birthday          time.Time         `json:"birthday" bson:"birthday"`
	Email             string            `json:"email" bson:"email"`
	DeliveryAddresses []DeliveryAddress `json:"delivery_addresses" bson:"delivery_addresses"`
	PhoneNumber       string            `json:"phone_number" bson:"phone_number"`
	ProfilePicture    string            `json:"profile_picture" bson:"profile_picture"`
}

type LoginRequest struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"-" bson:"password"`
}

type DeliveryAddress struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Address   string             `json:"address" bson:"address"`
	Phone     string             `json:"phone" bson:"phone"`
	IsDefault bool               `json:"is_default" bson:"is_default"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

const (
	ColName = "name"
)
