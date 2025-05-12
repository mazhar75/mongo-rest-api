package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mongo-rest-api/db"
	"mongo-rest-api/handlers"
)

func main() {
	// connect to Mongo and populate db.Collection
	db.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	addr := ":8080"
	fmt.Println("Server running at", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
