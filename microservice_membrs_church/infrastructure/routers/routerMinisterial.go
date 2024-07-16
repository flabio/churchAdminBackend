package routers

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/application/handler"
)

var (
	ministerial handler.IMinisterialController = handler.NewMinisterialController()
)

func NewMinisterialRouter(app *fiber.App) {

	apiMinisterial := app.Group("/api/ministerial")

	apiMinisterial.Get("/", func(c *fiber.Ctx) error {
		return ministerial.GetMinisterialFindAll(c)
	})
	apiMinisterial.Get("/:id", func(c *fiber.Ctx) error {
		return ministerial.GetMinisterialFindById(c)
	})
	apiMinisterial.Post("/", func(c *fiber.Ctx) error {
		return ministerial.CreateMinisterial(c)
	})
	apiMinisterial.Put("/:id", func(c *fiber.Ctx) error {
		return ministerial.UpdateMinisterial(c)
	})
	apiMinisterial.Delete("/:id", func(c *fiber.Ctx) error {
		return ministerial.DeleteMinisterial(c)
	})
}
