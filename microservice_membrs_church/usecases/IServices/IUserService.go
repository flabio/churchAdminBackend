package IServices

import "github.com/gofiber/fiber/v2"

type IUserService interface {
	GetFindUserAll(c *fiber.Ctx) error
	GetFindUserById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
