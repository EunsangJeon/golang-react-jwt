package main

import (
	"log"

	"github.com/EunsangJeon/golang-react-jwt/backend/db"
	"github.com/EunsangJeon/golang-react-jwt/backend/router"
	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	log.Println("Server starts with port 8080.")
	r.Run(":8080")
}
