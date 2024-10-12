package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 구조체 정의
type User struct {
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"_id"`
	Id        string             `json:"id" bson:"id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
}
