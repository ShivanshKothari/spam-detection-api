package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	PhoneNumber string             `bson:"phone_number"`
	Email       string             `bson:"email,omitempty"`
	Password    string             `bson:"password"`
}
