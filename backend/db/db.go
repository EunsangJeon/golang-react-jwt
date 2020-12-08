package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EunsangJeon/golang-react-jwt/backend/config"
)

// DB instance
var DB *sql.DB

// Connect to db with config
func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("Error: Could not open DB with config given")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with DB")
	}

	DB = db
}
