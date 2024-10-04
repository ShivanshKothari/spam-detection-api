package handlers

import (
	"context"
	"net/http"
	"spam-detection-api/db"
	"spam-detection-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func MarkSpam(c *gin.Context) {
	var spam models.Spam
	if err := c.ShouldBindJSON(&spam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if number is already marked as spam
	filter := bson.M{"phone_number": spam.PhoneNumber}
	var existingSpam models.Spam
	err := db.DB.Collection("spams").FindOne(context.TODO(), filter).Decode(&existingSpam)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Number is already marked as spam"})
		return
	}

	// Insert spam
	_, err = db.DB.Collection("spams").InsertOne(context.TODO(), spam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark number as spam"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Number marked as spam"})
}
