package main

import (
	"log"

	"github.com/a-andiadisasmita/graded-challenge-3-andiadisasmita/config"
)

func main() {
	// Connect to the database
	config.Connect()

	// Placeholder for further functionality
	log.Println("Application is running...")

	// Keep the application running
	select {}
}
