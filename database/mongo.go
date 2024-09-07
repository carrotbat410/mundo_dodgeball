package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB User 컬렉션 참조 변수
var UserCollection *mongo.Collection

// MongoDB 연결 설정
func ConnectMongo() {
	// MongoDB 클라이언트 설정
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// MongoDB 연결
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// User 컬렉션 선택
	UserCollection = client.Database("fiberdb").Collection("users")
	log.Println("Connected to MongoDB!")
}