package config

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Connect to PostgreSQL
	connString := "host=localhost user=postgres password=postgres123 dbname=banktest sslmode=disable"

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err.Error())
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error testing database connection:", err.Error())
	}

	fmt.Println("Successfully connected to PostgresSql database!")
}
