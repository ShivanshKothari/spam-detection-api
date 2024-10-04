package handlers

import (
	"context"
	"log"
	"net/http"
	"spam-detection-api/config"
	"spam-detection-api/db"
	"spam-detection-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// TODO: prevent from saving empty phone number and enter phone number from body
	var user models.User
	log.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(user)

	// Check if user already exists
	filter := bson.M{"phone_number": user.PhoneNumber}
	var existingUser models.User
	err := db.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&existingUser)
	log.Println(err)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "An account with the phone number already exists."})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Insert user
	res, err := db.DB.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": res.InsertedID})
}
func Login(c *gin.Context) {
	var login struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	filter := bson.M{"phone_number": login.PhoneNumber}
	var user models.User
	err := db.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Debugging: Print the hashed password and provided password

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
