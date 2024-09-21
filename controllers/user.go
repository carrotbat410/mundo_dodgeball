package controllers

import (
	"fiber_prac/models"
	"fiber_prac/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// 회원가입 핸들러 함수
func RegisterUser(c *fiber.Ctx) error {
	// 요청 바디에서 유저 정보 파싱
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 유저 등록 서비스 호출
	err := services.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 성공적으로 회원가입 완료
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// api := api.Group("/v1", middleware) 필요시 middleware + c.Next() 사용 가능

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("hello world")
// })

// app.Get("/sample/error", func(c *fiber.Ctx) error {
// 	return fiber.NewError(782, "Custom Error Message")
// })

func Login(c *fiber.Ctx) error {
	fmt.Println("로그인 컨트롤러 들어옴")

	var input models.LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// validation
	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
