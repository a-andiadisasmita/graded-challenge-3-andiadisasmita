package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// DB is the global database connection
var DB *pgx.Conn

// Connect initializes the database connection
func Connect() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get DATABASE_URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	// Connect to the database
	DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the connection
	var version string
	err = DB.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	log.Printf("Connected to database successfully. PostgreSQL version: %s", version)
}
