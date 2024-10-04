package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Spam struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
}
