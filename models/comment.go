package models

// Comment represents the comments table in the database
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	PostID  int    `json:"post_id"`
}
