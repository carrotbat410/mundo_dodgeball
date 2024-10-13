package controllers

import (
	"fiber_prac/models"
	"fiber_prac/services"
	"fiber_prac/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetRooms(c *fiber.Ctx) error {
	rooms, err := services.GetRooms()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch boards",
		})
	}

	// 게시글 목록을 JSON으로 응답
	return c.Status(fiber.StatusOK).JSON(rooms)
}

func CreateRoom(c *fiber.Ctx) error {
	var room models.Room

	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "body parsing error",
		})
	}
	userId := c.Locals("id").(string)
	username := c.Locals("username").(string)

	room.Title = username + "님의 방"
	room.Id = userId
	room.Username = username
	room.Content = ""
	room.Password = ""
	room.IsStart = false
	room.MaxParticipantNumber = 2
	room.NowParticipantNumber = 1
	room.ParticipantIdList = []string{userId}
	room.CreatedAt = utils.GetCurrentKoreaTime()
	room.UpdatedAt = utils.GetCurrentKoreaTime()

	if err := services.CreateRoom(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//TODO 소켓 해당 방 입장 로직
	fmt.Println("웹소켓 - 해당 방에 입장 로직 필요")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})

}
