package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Episode struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
}
