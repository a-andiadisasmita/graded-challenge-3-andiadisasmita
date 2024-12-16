package controllers

import (
	"net/http"
	"regexp"

	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/config"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/models"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// CreatePost handles creating a new post
func CreatePost(c *gin.Context) {
	// Extract user ID from the JWT token
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate image URL
	urlPattern := `^https?:\/\/[^\s]+$`
	matched, _ := regexp.MatchString(urlPattern, post.ImageURL)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL"})
		return
	}

	// Use random joke if content is not provided
	if post.Content == "" {
		client := resty.New()
		resp, err := client.R().
			SetHeader("X-Api-Key", utils.GetJokesAPIKey()).
			Get("https://api.api-ninjas.com/v1/jokes")
		if err != nil || resp.StatusCode() != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch joke content"})
			return
		}

		post.Content = string(resp.Body())
	}

	// Insert post into the database
	query := `INSERT INTO posts (content, image_url, user_id) VALUES ($1, $2, $3) RETURNING id`
	err = config.DB.QueryRow(c, query, post.Content, post.ImageURL, userID).Scan(&post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// GetPosts retrieves all posts
func GetPosts(c *gin.Context) {
	rows, err := config.DB.Query(c, `SELECT id, content, image_url, user_id FROM posts`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Content, &post.ImageURL, &post.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse posts"})
			return
		}
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
