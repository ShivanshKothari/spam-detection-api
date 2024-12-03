package main

import (
	"spam-detection-api/db"
	"spam-detection-api/handlers"
	"spam-detection-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize router
	r := gin.Default()

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Next.js default port
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	r.Use(cors.New(corsConfig))

	// Initialize database connection
	db.Init()

	// Public routes
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/mark-spam", handlers.MarkSpam)
		protected.GET("/searchbyname", handlers.SearchByName)
		protected.GET("/searchbynumber", handlers.SearchByPhoneNumber)
	}

	// Start server
	r.Run(":8080")
}
