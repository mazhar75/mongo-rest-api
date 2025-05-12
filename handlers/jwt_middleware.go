package handlers

import (
  "context"
  "net/http"
  "strings"

  "github.com/golang-jwt/jwt/v5"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"

  "mongo-rest-api/db"
  "mongo-rest-api/models"
)

//var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// AuthMiddleware protects routes by validating the JWT
func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    auth := r.Header.Get("Authorization")
    if !strings.HasPrefix(auth, "Bearer ") {
      http.Error(w, "missing token", http.StatusUnauthorized)
      return
    }
    tokenStr := strings.TrimPrefix(auth, "Bearer ")

    token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
      return jwtKey, nil
    })
    if err != nil || !token.Valid {
      http.Error(w, "invalid token", http.StatusUnauthorized)
      return
    }
    claims := token.Claims.(*jwt.RegisteredClaims)

    // optional: load the user into context
    id, _ := primitive.ObjectIDFromHex(claims.Subject)
    var u models.User
    _ = db.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&u)
    ctx := context.WithValue(r.Context(), "user", u)

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
