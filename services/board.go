package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"fiber_prac/database"
	"fiber_prac/models"

	"go.mongodb.org/mongo-driver/bson"
)

// var boardCollection *mongo.Collection

// func InitDatabase(client *mongo.Client) {
// 	boardCollection = client.Database("mydb").Collection("boards")
// }

func GetBoards() ([]models.Board, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var boards []models.Board

	// MongoDB에서 모든 게시글을 찾습니다.
	// cursor, err := boardCollection.Find(ctx, bson.M{})
	cursor, err := database.BoardCollection.Find(ctx, bson.M{})

	fmt.Println("cursor:", cursor)

	if err != nil {
		log.Println("Error fetching boards from database:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// 결과를 배열로 변환
	if err = cursor.All(ctx, &boards); err != nil {
		log.Println("Error parsing boards:", err)
		return nil, err
	}

	return boards, nil
}

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
