package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Password    string             `bson:"password"`
}
