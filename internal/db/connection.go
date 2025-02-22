package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // driver MySQL
)

var db *sql.DB

// function InitDB inicialize the connection with MySQL
func InitDB() {
	var err error

	// constructing the connection database string in .env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// verifying error in definition variable
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Error: One or more database enviroment variables were not definided")
	}

	// taking the environment connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// open the connection
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
	}

	// verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Databse inaccessible: %v", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
}

// GetDB return the instance database
func GetDB() *sql.DB {
	return db
}
