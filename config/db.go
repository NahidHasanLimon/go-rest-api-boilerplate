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
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBSSLMode,
	)

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