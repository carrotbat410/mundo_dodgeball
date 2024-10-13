package services

import (
	"context"
	"errors"
	"log"
	"time"

	"fiber_prac/database"
	"fiber_prac/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var rooms []models.Room

	cursor, err := database.RoomCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Println("Error fetching boards from database:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func CreateRoom(room models.Room) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := database.RoomCollection.InsertOne(ctx, room)

	if err != nil {
		return errors.New("Failed to create room")
	}

	return nil
}
