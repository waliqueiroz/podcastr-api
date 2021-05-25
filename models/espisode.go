package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Episode struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Slug        string             `json:"slug" bson:"slug"`
	Title       string             `json:"title" bson:"title"`
	Members     string             `json:"members" bson:"members"`
	PublishedAt time.Time          `json:"published_at" bson:"published_at"`
	Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
	Description string             `json:"description" bson:"description"`
	File        File               `json:"file" bson:"file"`
}

type File struct {
	URL      string `json:"url" bson:"url"`
	Type     string `json:"type" bson:"type"`
	Duration int64  `json:"duration" bson:"duration"`
}
