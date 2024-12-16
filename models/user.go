package models

// User represents the user table in the database
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	Age      int    `json:"age"`
}
