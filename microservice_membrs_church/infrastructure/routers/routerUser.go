package routers

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/application/handler"
)

var (
	user handler.IUserController = handler.NewUserController()
)

func NewUserRouter(app *fiber.App) {
	api := app.Group("/api/user")
	api.Get("/", func(c *fiber.Ctx) error {
		return user.GetFindUserAll(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return user.GetFindUserById(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return user.CreateUser(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return user.UpdateUser(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return user.DeleteUser(c)
	})
}
