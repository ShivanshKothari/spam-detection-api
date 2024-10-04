package main

import (
	"spam-detection-api/db"
	"spam-detection-api/handlers"
	"spam-detection-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/mark-spam", handlers.MarkSpam)
		auth.GET("/searchbyname", handlers.SearchByName)
		auth.GET("/searchbynumber", handlers.SearchByPhoneNumber)
	}

	r.Run(":8080")
}
