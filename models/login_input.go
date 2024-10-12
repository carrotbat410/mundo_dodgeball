package models

type LoginInput struct {
	Id       string `json:"id" bson:"id"`
	Password string `json:"password" bson:"password"`
}
