package main

import (
	"log"

	"foodos-backend/internal/server"
)

func main() {
	srv := server.New()
	log.Println("🚀 FoodOS backend running on :8080")
	srv.Run(":8080")
}
