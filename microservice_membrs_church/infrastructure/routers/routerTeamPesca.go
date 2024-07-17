package routers

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/application/handler"
)

var (
	teamPescaController handler.ITeamPescaController = handler.NewTeamPescaController()
)

func NewTeamPescaRouter(app *fiber.App) {
	apiTeamPesca := app.Group("/api/teampesca")
	apiTeamPesca.Get("/", func(c *fiber.Ctx) error {
		return teamPescaController.GetTeamPescaFindAll(c)
	})
	apiTeamPesca.Get("/:id", func(c *fiber.Ctx) error {
		return teamPescaController.GetTeamPescaFindById(c)
	})
	apiTeamPesca.Post("/", func(c *fiber.Ctx) error {
		return teamPescaController.CreateTeamPesca(c)
	})
	apiTeamPesca.Put("/:id", func(c *fiber.Ctx) error {
		return teamPescaController.UpdateTeamPesca(c)
	})
	apiTeamPesca.Delete("/:id", func(c *fiber.Ctx) error {
		return teamPescaController.DeleteTeamPesca(c)
	})
}
