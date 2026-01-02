package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ failed to connect postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ postgres ping failed: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return db
}
