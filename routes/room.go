package routes

import (
	"fiber_prac/controllers"
	"fiber_prac/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RoomRoutes(app *fiber.App) {
	group := app.Group("/room")

	group.Get("/", middlewares.AuthRequired, controllers.GetRooms)
	group.Post("/", middlewares.AuthRequired, controllers.CreateRoom)

	//TODO 1 - 입장하기 API
	//중복 입장 못하도록

	//TODO 2 - 방 나가기 api만들기 (현재 인원수가 0이 되면 안보이도록? <- 복잡해질듯)
	//더 간단하게 방장이 나가면 방 폭파되도록(남은 유저들에게는 알럿뜨게) 하는게 깔끔할듯

}
