package main

import (
	"log"
	"os"

	"fiber_prac/database"
	"fiber_prac/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// 환경변수 가져오기
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 서버 포트 설정
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":3000"
	}

	// MongoDB 연결
	database.ConnectMongo()

	// Fiber 인스턴스 생성
	app := fiber.New()

	// 라우트 설정
	setupRoutes(app)

	// 서버 시작
	log.Fatal(app.Listen(serverPort))
}

// 라우트 설정 함수
func setupRoutes(app *fiber.App) {
	routes.AuthRoutes(app)
	routes.BoardRoutes(app)
	routes.RoomRoutes(app)
}
