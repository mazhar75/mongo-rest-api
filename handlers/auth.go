package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
    "log"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"mongo-rest-api/db"
	"mongo-rest-api/models"
)

// jwtKey is read once from the environment
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// standard JSON error payload
func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// RegisterUser handles POST /register
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "could not hash password")
		return
	}
	u.Password = string(hashed)

	// insert user
	if _, err := db.Collection.InsertOne(context.TODO(), u); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"msg": "user created"})
}

// LoginUser handles POST /login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
     log.Printf("[Login] Attempting login for email=%q password=%q\n", creds.Email, creds.Password)
	// fetch user by email
	var u models.User
	if err := db.Collection.FindOne(context.TODO(), bson.M{"email": creds.Email}).Decode(&u); err != nil {
		writeJSONError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}
    log.Printf("[Login] Found user. Stored bcrypt hash: %q\n", u.Password)
	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(creds.Password)); err != nil {
		writeJSONError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}
	log.Printf("[Login] Password match for %q! Generating token...\n", creds.Email)

	// create JWT token
	claims := &jwt.RegisteredClaims{
		Subject:   u.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "could not create token")
		return
	}

	// respond with JSON token
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
