package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	BoardId   primitive.ObjectID `json:"board_id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Id        string             `json:"id" bson:"id"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
