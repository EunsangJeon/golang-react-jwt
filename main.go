package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Product is struture for product
type Product struct {
	ID   int
	Name string
	Slug string
	Desc string
}

var products = []Product{
	{ID: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Desc: "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{ID: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Desc: "Explore the depths of the sea in this one of a kind underwater experience"},
	{ID: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Desc: "Go back 65 million years in the past and ride a T-Rex"},
	{ID: 4, Name: "Cars VR", Slug: "cars-vr", Desc: "Get behind the wheel of the fastest cars in the world."},
	{ID: 5, Name: "Robin Hood", Slug: "robin-hood", Desc: "Pick up the bow and arrow and master the art of archery"},
	{ID: 6, Name: "Real World VR", Slug: "real-world-vr", Desc: "Explore the seven wonders of the world in VR"},
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/products", productsHandler).Methods("GET")
	r.Handle("/products/{Slug}/feedback", addFeedbackHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running."))
})

var productsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
})

var addFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var p Product
	vars := mux.Vars(r)
	Slug := vars["Slug"]

	for _, product := range products {
		if product.Slug == Slug {
			p = product
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if p.Slug != "" {
		payload, _ := json.Marshal(p)
		w.Write(payload)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})
