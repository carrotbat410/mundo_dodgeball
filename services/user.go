package services

import (
	"context"
	"errors"
	"time"

	"fiber_prac/database"
	"fiber_prac/models"
	"fiber_prac/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// 유저 등록 서비스
func RegisterUser(user models.User) error {

	// 중복 이메일 확인
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count1, err1 := database.UserCollection.CountDocuments(ctx, bson.M{"id": user.Id})
	if err1 != nil {
		return errors.New("Database error")
	}
	if count1 > 0 {
		return errors.New("id already exists")
	}

	count2, err2 := database.UserCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err2 != nil {
		return errors.New("Database error")
	}
	if count2 > 0 {
		return errors.New("username already exists")
	}

	// 유저 데이터 삽입
	user.CreatedAt = utils.GetCurrentKoreaTime()

	_, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return errors.New("Failed to register user")
	}

	return nil
}
