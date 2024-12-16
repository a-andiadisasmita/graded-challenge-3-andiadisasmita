package routes

import (
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/users/register", controllers.RegisterUser) // Register a new user
	router.POST("/users/login", controllers.LoginUser)       // User login

	// Post routes
	router.POST("/posts", controllers.CreatePost) // Create a new post
	router.GET("/posts", controllers.GetPosts)    // Retrieve all posts

	// Comment routes
	router.POST("/comments", controllers.CreateComment) // Add a comment to a post

	// Serve Swagger documentation
	router.Static("/swagger", "./swagger")

	return router
}
