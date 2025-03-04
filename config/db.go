package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
var DB *sql.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=my_new_database sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection not established:", err)
	}
	fmt.Println("Connected to PostgreSQL database successfully!")

}