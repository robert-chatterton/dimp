package route

import (
	"main/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", controller.Login)
	app.Delete("/api/login", controller.Logout)
	app.Post("/api/register", controller.Register)

	app.Get("/api/user", controller.GetUser)
}
