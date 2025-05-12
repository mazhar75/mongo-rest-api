package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user document in MongoDB
type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}
