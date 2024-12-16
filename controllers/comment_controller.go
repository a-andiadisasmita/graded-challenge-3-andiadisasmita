package controllers

import (
	"net/http"

	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/config"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/models"
	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/utils"
	"github.com/gin-gonic/gin"
)

// CreateComment handles creating a new comment
func CreateComment(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Insert comment into database
	query := `INSERT INTO comments (content, user_id, post_id) VALUES ($1, $2, $3) RETURNING id`
	err = config.DB.QueryRow(c, query, comment.Content, userID, comment.PostID).Scan(&comment.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"comment": comment,
	})
}
