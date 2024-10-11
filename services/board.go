package services

import (
	"context"
	"errors"
	"time"

	"fiber_prac/database"
	"fiber_prac/models"
)

// func RegisterUser(user models.User) error {

// 	if user.Username == "" || user.Email == "" || user.Password == "" {
// 		return errors.New("All fields are required")
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
// 	if err != nil {
// 		return errors.New("Database error")
// 	}
// 	if count > 0 {
// 		return errors.New("Email already exists")
// 	}

// 	user.CreatedAt = utils.GetCurrentKoreaTime()

// 	_, err = database.UserCollection.InsertOne(ctx, user)
// 	if err != nil {
// 		return errors.New("Failed to register user")
// 	}

// 	return nil
// }

func CreateBoard(board models.Board) error {
	if board.Content == "" || board.Title == "" {
		return errors.New("All fields are required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.BoardCollection.InsertOne(ctx, board)

	if err != nil {
		return errors.New("Failed to create board")
	}

	return nil
}
