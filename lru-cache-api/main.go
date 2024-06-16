package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"lru-cache-api/api"
	// Using gorilla/mux for routing
)

func main() {
	router := api.InitializeRouter()

	// CORS configuration
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Adjust to your frontend URL
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})

	// CORS middleware
	corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	log.Println("Server started on :8005")
	log.Fatal(http.ListenAndServe(":8005", corsHandler))
}
