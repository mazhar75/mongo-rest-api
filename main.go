package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"mongo-rest-api/db"
	"mongo-rest-api/handlers"
)

// init loads environment variables from .env before anything else
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — relying on OS environment")
	}
}

func main() {
	// Ensure JWT secret is set
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET not set")
	}

	// Connect to MongoDB
	db.Connect()

	// Create router
	r := mux.NewRouter()

	// Public endpoints
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Protected endpoints under /users
	api := r.PathPrefix("/users").Subrouter()
	api.Use(handlers.AuthMiddleware)
	api.HandleFunc("", handlers.CreateUser).Methods("POST")
	api.HandleFunc("", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
	api.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// Start server
	addr := ":8080"
	fmt.Println("Server running at", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
