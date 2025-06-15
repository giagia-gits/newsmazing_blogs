package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Go API!",
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")

	// Define allowed CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://3.22.70.110:3000"}), // or specify frontend origin
		handlers.AllowedOrigins([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Wrap the router with the CORS middleware
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsOptions(r)))
}
