package controllers

import (
	"fiber_prac/models"
	"fiber_prac/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// func GetBaords(c *fiber.Ctx) error {
// 	// 요청 바디에서 유저 정보 파싱
// 	var boards models.Board
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	// 유저 등록 서비스 호출
// 	err := services.RegisterUser(user)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	// 성공적으로 회원가입 완료
// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 	"message": "User registered successfully",
// })
// }

func CreateBoard(c *fiber.Ctx) error {
	var board models.Board

	if err := c.BodyParser(&board); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body parsing error",
		})
	}

	fmt.Println("board info:", board)

	if err := services.CreateBoard(board); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})

}
