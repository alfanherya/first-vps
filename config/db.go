package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// init DB initialize the database connection
func InitDb() {
	// Load environment variables
	InitEnv()

	// Get database connection string from environtment variable
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_NAME", "postgres"))

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	fmt.Println("Successfully connected to the database!")

}
