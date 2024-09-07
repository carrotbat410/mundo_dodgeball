package routes

import (
	"fiber_prac/controllers"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes는 /auth 그룹 경로를 설정합니다.
func AuthRoutes(app *fiber.App) {
	// /auth 그룹 생성
	authGroup := app.Group("/auth")

	// /auth/register 경로 정의
	authGroup.Post("/register", controllers.RegisterUser)
}

// // 라우트 설정 함수
// func setupRoutes(app *fiber.App) {
// 	app.Post("/register", controllers.RegisterUser)
// }

// api := api.Group("/v1", middleware) 필요시 middleware + c.Next() 사용 가능

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("hello world")
// })

// app.Get("/sample/error", func(c *fiber.Ctx) error {
// 	return fiber.NewError(782, "Custom Error Message")
// })

// //* 매개변수
// app.Get("/sample/:name", func(c *fiber.Ctx) error {
// 	return c.SendString("name: " + c.Params("name"))
// })
