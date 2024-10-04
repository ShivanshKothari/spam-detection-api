package handlers

import (
	"context"
	"net/http"
	"spam-detection-api/db"
	"spam-detection-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		return
	}

	// Search in contacts
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: "^" + name, Options: "i"}}}
	cursor, err := db.DB.Collection("contacts").Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search contacts"})
		return
	}
	var contacts []models.Contact
	if err = cursor.All(context.TODO(), &contacts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode contacts"})
		return
	}

	// Search in users
	filter = bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: "^" + name, Options: "i"}}}
	cursor, err = db.DB.Collection("users").Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode users"})
		return
	}

	results := make([]gin.H, 0)
	for _, contact := range contacts {
		results = append(results, gin.H{
			"name":            contact.Name,
			"phone_number":    contact.PhoneNumber,
			"spam_likelihood": getSpamLikelihood(contact.PhoneNumber),
		})
	}
	for _, user := range users {
		results = append(results, gin.H{
			"name":            user.Name,
			"phone_number":    user.PhoneNumber,
			"spam_likelihood": getSpamLikelihood(user.PhoneNumber),
		})
	}

	c.JSON(http.StatusOK, results)
}

func SearchByPhoneNumber(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number query parameter is required"})
		return
	}

	// Search in users
	filter := bson.M{"phone_number": phoneNumber}
	var user models.User
	err := db.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"name":            user.Name,
			"phone_number":    user.PhoneNumber,
			"spam_likelihood": getSpamLikelihood(user.PhoneNumber),
		})
		return
	}

	// Search in contacts
	filter = bson.M{"phone_number": phoneNumber}
	cursor, err := db.DB.Collection("contacts").Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search contacts"})
		return
	}
	var contacts []models.Contact
	if err = cursor.All(context.TODO(), &contacts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode contacts"})
		return
	}

	results := make([]gin.H, 0)
	for _, contact := range contacts {
		results = append(results, gin.H{
			"name":            contact.Name,
			"phone_number":    contact.PhoneNumber,
			"spam_likelihood": getSpamLikelihood(contact.PhoneNumber),
		})
	}

	c.JSON(http.StatusOK, results)
}

func getSpamLikelihood(phoneNumber string) int {
	filter := bson.M{"phone_number": phoneNumber}
	count, err := db.DB.Collection("spams").CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0
	}
	return int(count)
}
