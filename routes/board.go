package routes

import (
	"fiber_prac/controllers"

	"github.com/gofiber/fiber/v2"
)

func BoardRoutes(app *fiber.App) {
	group := app.Group("/board")

	// group.Get("/boards", controllers.GetBaords)
	group.Post("/boards", controllers.CreateBoard)
	// authGroup.Post("/register", controllers.RegisterUser)

}
