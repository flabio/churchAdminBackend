package routers

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/application/handler"
)

var (
	church handler.IChurchController = handler.NewChurchController()
)

func NewChurchRouter(app *fiber.App) {

	api := app.Group("/api/church")
	api.Get("/", func(c *fiber.Ctx) error {
		return church.GetChurchFindAll(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return church.GetChurchFindById(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return church.CreateChurch(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return church.UpdateChurch(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return church.DeleteChurch(c)
	})

}
