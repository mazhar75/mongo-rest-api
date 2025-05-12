package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo-rest-api/db"
	"mongo-rest-api/models"
)

// CreateUser handles POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := db.Collection.InsertOne(context.TODO(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetUsers handles GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cursor, err := db.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	for cursor.Next(context.TODO()) {
		var u models.User
		cursor.Decode(&u)
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

// GetUser handles GET /users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var u models.User
	if err := db.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&u); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

// UpdateUser handles PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	update := bson.M{"$set": bson.M{"name": u.Name, "email": u.Email}}
	if _, err := db.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("User updated")
}

// DeleteUser handles DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	if _, err := db.Collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("User deleted")
}
