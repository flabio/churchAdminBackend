package handler

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/services"
)

type IMinisterialController interface {
	GetMinisterialFindAll(c *fiber.Ctx) error
	GetMinisterialFindById(c *fiber.Ctx)error
	CreateMinisterial(c *fiber.Ctx)error
	UpdateMinisterial(c *fiber.Ctx)error
	DeleteMinisterial(c *fiber.Ctx)error
}

type ministerialController struct {
	ministerial IServices.IMinisterialService
}

func NewMinisterialController() IMinisterialController {
	return &ministerialController{
		ministerial: services.NewMinisterialService(),
	}
}

func (m *ministerialController) GetMinisterialFindAll(c *fiber.Ctx)error {
	return m.ministerial.GetMinisterialFindAll(c)
}
func (m *ministerialController) GetMinisterialFindById(c *fiber.Ctx) error{
	return m.ministerial.GetMinisterialFindById(c)
}
func (m *ministerialController) CreateMinisterial(c *fiber.Ctx) error{
	return m.ministerial.CreateMinisterial(c)
}
func (m *ministerialController) UpdateMinisterial(c *fiber.Ctx) error{
	return m.ministerial.UpdateMinisterial(c)
}
func (m *ministerialController) DeleteMinisterial(c *fiber.Ctx) error{
	return m.ministerial.DeleteMinisterial(c)
}
