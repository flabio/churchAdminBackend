package handler

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/services"
)

type RolController interface {
	GetFindAll(c *fiber.Ctx) error
	GetFindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type rolController struct {
	rol IServices.IRolService
}

func NewRolController() RolController {

	return &rolController{
		rol: services.NewRolService(),
	}
}

func (r *rolController) GetFindAll(c *fiber.Ctx) error {
	return r.rol.GetFindAll(c)

}

func (r *rolController) GetFindById(c *fiber.Ctx) error {
	return r.rol.GetFindById(c)

}

func (r *rolController) Create(c *fiber.Ctx) error {
	return r.rol.Create(c)
}
func (r *rolController) Update(c *fiber.Ctx) error{
	return r.rol.Update(c)
}

func (r *rolController) Delete(c *fiber.Ctx) error{
	return r.rol.Delete(c)
}
