package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ErrorRecord represents an error record in the database
type ErrorRecord struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Data string             `bson:"data"`
}
