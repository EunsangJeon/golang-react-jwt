package main

import (
	_ "database/sql"

	"github.com/EunsangJeon/golang-react-jwt/backend/db"
	"github.com/EunsangJeon/golang-react-jwt/backend/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8080
	r.Run(":8080")
}
