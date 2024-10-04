package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Name        string             `bson:"name"`
	PhoneNumber string             `bson:"phone_number"`
}
