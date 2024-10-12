package routes

import (
	"fiber_prac/controllers"
	"fiber_prac/middlewares"

	"github.com/gofiber/fiber/v2"
)

func BoardRoutes(app *fiber.App) {
	group := app.Group("/board")

	group.Get("/", controllers.GetBoards)
	group.Post("/", middlewares.AuthRequired, controllers.CreateBoard)
	// authGroup.Post("/register", controllers.RegisterUser)

}
