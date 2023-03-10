package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "user-rest-api/controllers"
)

func HandleRequests() {
	app := fiber.New()
	app.Get("/", controller.FindAll)
	app.Get("/:code", controller.FindByCode)
	app.Post("/", controller.CreateUSer)
	app.Listen(":8080")
}
