package models

import "time"

type Board struct {
	BoardId   string    `json:"board_id,omitempty" bson:"board_id,omitempty"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
