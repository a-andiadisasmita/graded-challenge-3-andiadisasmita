package models

// Post represents the posts table in the database
type Post struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
	UserID   int    `json:"user_id"`
}
