package handler

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/services"
)

type ITeamPescaController interface {
	GetTeamPescaFindAll(c *fiber.Ctx) error
	GetTeamPescaFindById(c *fiber.Ctx) error
	CreateTeamPesca(c *fiber.Ctx) error
	UpdateTeamPesca(c *fiber.Ctx) error
	DeleteTeamPesca(c *fiber.Ctx) error
}
type teamPescaController struct {
	teamPesca IServices.ITeamPescaService
}

func NewTeamPescaController() ITeamPescaController {
	return &teamPescaController{
		teamPesca: services.NewTeamPescaService(),
	}
}

func (t *teamPescaController) GetTeamPescaFindAll(c *fiber.Ctx) error {
	return t.teamPesca.GetTeamPescaFindAll(c)
}

func (t *teamPescaController) GetTeamPescaFindById(c *fiber.Ctx) error {
	return t.teamPesca.GetTeamPescaFindById(c)
}
func (t *teamPescaController) CreateTeamPesca(c *fiber.Ctx) error {
	return t.teamPesca.CreateTeamPesca(c)
}

func (t *teamPescaController) UpdateTeamPesca(c *fiber.Ctx) error {
	return t.teamPesca.UpdateTeamPesca(c)
}

func (t *teamPescaController) DeleteTeamPesca(c *fiber.Ctx) error {
	return t.teamPesca.DeleteTeamPesca(c)
}
