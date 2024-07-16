package handler

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/services"
)

type IChurchController interface {
	GetChurchFindAll(c *fiber.Ctx) error
	GetChurchFindById(c *fiber.Ctx) error
	CreateChurch(c *fiber.Ctx) error
	UpdateChurch(c *fiber.Ctx) error
	DeleteChurch(c *fiber.Ctx) error
}

type churchController struct {
	churchService IServices.IChurchService
}

func NewChurchController() IChurchController {
	return &churchController{
		churchService: services.NewChurchService(),
	}
}

func (controller *churchController) GetChurchFindAll(c *fiber.Ctx) error {
	return controller.churchService.GetChurchFindAll(c)
}
func (controller *churchController) GetChurchFindById(c *fiber.Ctx) error {
	return controller.churchService.GetChurchFindById(c)
}
func (controller *churchController) CreateChurch(c *fiber.Ctx) error {
	return controller.churchService.CreateChurch(c)
}
func (controller *churchController) UpdateChurch(c *fiber.Ctx) error {
	return controller.churchService.UpdateChurch(c)
}
func (controller *churchController) DeleteChurch(c *fiber.Ctx) error {
	return controller.churchService.DeleteChurch(c)
}
