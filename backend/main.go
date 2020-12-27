package main

import (
	"log"
	"net/http"

	"github.com/EunsangJeon/golang-react-jwt/backend/config"
	"github.com/EunsangJeon/golang-react-jwt/backend/controller"
	"github.com/EunsangJeon/golang-react-jwt/backend/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.Create).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/login", controller.Login).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/session", controller.Session).Methods(http.MethodGet, http.MethodOptions)

	log.Println("Server starts with port 8080")
	log.Fatalln(
		http.ListenAndServe(
			":8080",
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST"}),
				handlers.AllowedOrigins([]string{config.ClientURL}),
			)(r),
		),
	)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
