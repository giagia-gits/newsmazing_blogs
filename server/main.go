package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go API!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")

	// Define allowed CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // or specify frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Wrap the router with the CORS middleware
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsOptions(r)))
}
