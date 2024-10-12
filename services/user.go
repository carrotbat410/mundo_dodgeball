package services

import (
	"context"
	"errors"
	"time"

	"fiber_prac/database"
	"fiber_prac/models"
	"fiber_prac/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// 유저 데이터 삽입
	user.Password = string(hashedPassword)
	user.CreatedAt = utils.GetCurrentKoreaTime()

	_, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return errors.New("Failed to register user")
	}

	return nil
}

func Login(id string, password string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var findUser models.User

	err := database.UserCollection.FindOne(ctx, bson.M{"id": id}).Decode(&findUser)

	if err == mongo.ErrNoDocuments {
		return findUser, errors.New("User not found")
	} else if err != nil {
		// 다른 오류 처리
		// return findUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if err != nil {
		return findUser, errors.New("User not found")
	}

	return findUser, nil
}
