package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

// InitDB initializes the PostgreSQL database connection
func InitDB() error {
	var err error

	// Get connection string from environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL environment variable is required. Please set it in your .env file.")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(0)

	// Test connection
	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("✅ Database connection established")
	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
