package db

import (
  "context"
  "fmt"
  "log"
  "os"
  "github.com/joho/godotenv"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func Connect() {
  // load .env (only in non-production or local dev)
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found, falling back to OS env vars")
  }

  uri := os.Getenv("MONGO_URI")
  if uri == "" {
    log.Fatal("MONGO_URI must be set")
  }

  clientOpts := options.Client().ApplyURI(uri)
  client, err := mongo.Connect(context.TODO(), clientOpts)
  if err != nil {
    log.Fatalf("Mongo connect error: %v", err)
  }
  if err := client.Ping(context.TODO(), nil); err != nil {
    log.Fatalf("Mongo ping error: %v", err)
  }

  fmt.Println("✅ Connected to MongoDB")
  Collection = client.Database("testdb").Collection("users")
}
