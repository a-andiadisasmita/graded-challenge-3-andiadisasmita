package controllers

import (
	"log"
	"net/http"

	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/config"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/models"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/utils"
	"github.com/gin-gonic/gin"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User

	// Bind JSON request to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Insert user into database
	query := `INSERT INTO users (full_name, email, username, password, age) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := config.DB.QueryRow(c, query, user.FullName, user.Email, user.Username, user.Password, user.Age).Scan(&user.ID)
	if err != nil {
		log.Println("Error inserting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
			"username":  user.Username,
			"age":       user.Age,
		},
	})
}

// LoginUser handles user login and returns a JWT token
func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Verify user credentials
	var user models.User
	query := `SELECT id, full_name, email, username, age FROM users WHERE email = $1 AND password = $2`
	err := config.DB.QueryRow(c, query, input.Email, input.Password).Scan(&user.ID, &user.FullName, &user.Email, &user.Username, &user.Age)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error generating JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
