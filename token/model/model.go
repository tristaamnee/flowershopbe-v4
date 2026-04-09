package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id" pg:"type:uuid,pk"`
	CreationDate time.Time          `json:"creationDate" bson:"creationDate"`
	UserID       primitive.ObjectID `json:"userId" bson:"userId"`
	Activities   []*TokenActivity   `json:"activities" bson:"activities"`
	CreateAt     time.Time          `json:"createAt" bson:"createAt"`
	UpdateAt     time.Time          `json:"updateAt" bson:"updateAt"`
}

type TokenActivity struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id" pg:"type:uuid,pk"`
	TokenID  primitive.ObjectID `json:"tokenId" bson:"tokenId"`
	Endpoint string             `json:"endpoint" bson:"endpoint"`
	CreateAt time.Time          `json:"createAt" bson:"createAt"`
}
