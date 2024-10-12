package controllers

import (
	"fiber_prac/models"
	"fiber_prac/services"
	"fiber_prac/utils"

	"github.com/gofiber/fiber/v2"
)

// func GetBoards(c *fiber.Ctx) error {
// 	var boards []models.Board
// }

func GetBoards(c *fiber.Ctx) error {
	// 서비스 계층에서 모든 게시글 가져오기
	boards, err := services.GetBoards()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch boards",
		})
	}

	// 게시글 목록을 JSON으로 응답
	return c.Status(fiber.StatusOK).JSON(boards)
}

func CreateBoard(c *fiber.Ctx) error {
	var board models.Board

	if err := c.BodyParser(&board); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body parsing error",
		})
	}

	board.CreatedAt = utils.GetCurrentKoreaTime()
	board.UpdatedAt = utils.GetCurrentKoreaTime()

	if err := services.CreateBoard(board); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})

}
