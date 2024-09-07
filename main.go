package main

import (
	"log"

	"fiber_prac/database"
	"fiber_prac/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// MongoDB 연결
	database.ConnectMongo()

	// Fiber 인스턴스 생성
	app := fiber.New()

	// 라우트 설정
	setupRoutes(app)

	// 서버 시작
	log.Fatal(app.Listen(":3000"))
}

// 라우트 설정 함수
func setupRoutes(app *fiber.App) {
	// auth 관련 라우트 설정
	routes.AuthRoutes(app)
}
