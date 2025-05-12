package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "mongo-rest-api/db"
  "mongo-rest-api/handlers"
)

func main() {
  // 1. connect to MongoDB
  db.Connect()

  // 2. set up router
  r := mux.NewRouter()
  r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
  r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
  r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
  r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
  r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

  // 3. configure CORS
  c := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},          // allow all, or list your domains
    AllowedMethods:   []string{"GET","POST","PUT","DELETE","OPTIONS"},
    AllowedHeaders:   []string{"Authorization","Content-Type"},
    AllowCredentials: true,
  })

  // 4. wrap router with CORS middleware
  handler := c.Handler(r)

  addr := ":8080"
  fmt.Println("Server running at", addr)
  log.Fatal(http.ListenAndServe(addr, handler))
}
