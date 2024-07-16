package handler

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/services"
)

type IUserController interface {
	GetFindUserAll(c *fiber.Ctx) error
	GetFindUserById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserController struct {
	user IServices.IUserService
}

func NewUserController() IUserController {

	return &UserController{
		user: services.NewUserService(),
	}
}

func (r *UserController) GetFindUserAll(c *fiber.Ctx) error {
	return r.user.GetFindUserAll(c)

}

func (r *UserController) GetFindUserById(c *fiber.Ctx) error {
	return r.user.GetFindUserById(c)

}

func (r *UserController) CreateUser(c *fiber.Ctx) error {
	return r.user.CreateUser(c)
}
func (r *UserController) UpdateUser(c *fiber.Ctx) error{
	return r.user.UpdateUser(c)
}

func (r *UserController) DeleteUser(c *fiber.Ctx) error{
	return r.user.DeleteUser(c)
}
