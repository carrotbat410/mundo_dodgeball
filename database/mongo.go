package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB User 컬렉션 참조 변수
var UserCollection *mongo.Collection
var BoardCollection *mongo.Collection

// MongoDB 연결 설정
func ConnectMongo() {
	// 환경 변수 가져오기
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// MongoDB 클라이언트 설정
	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		mongoURL = "localhost:27017"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoURL))
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
	BoardCollection = client.Database("fiberdb").Collection("boards")

	log.Println("Connected to MongoDB!")
}
