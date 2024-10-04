package main

import (
	"context"
	"fmt"
	"log"
	"spam-detection-api/config"
	"spam-detection-api/models"
	"spam-detection-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.DbName)

	// Populate Users
	usersCollection := db.Collection("users")
	for i := 0; i < 100; i++ {
		user := models.User{
			Name:        fmt.Sprintf("User%d", i),
			PhoneNumber: fmt.Sprintf("123456789%d", i),
			Email:       fmt.Sprintf("user%d@example.com", i),
			Password:    utils.GenerateRandomString(10),
		}
		user.Password, _ = utils.HashPassword(user.Password)
		_, err := usersCollection.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Populate Contacts
	contactsCollection := db.Collection("contacts")
	for i := 0; i < 100; i++ {
		contact := models.Contact{
			UserID:      primitive.NewObjectID(),
			Name:        fmt.Sprintf("Contact%d", i),
			PhoneNumber: fmt.Sprintf("987654321%d", i),
		}
		_, err := contactsCollection.InsertOne(ctx, contact)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Populate Spams
	spamsCollection := db.Collection("spams")
	for i := 0; i < 100; i++ {
		spam := models.Spam{
			PhoneNumber: fmt.Sprintf("112233445%d", i),
		}
		_, err := spamsCollection.InsertOne(ctx, spam)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data population completed successfully.")
}
