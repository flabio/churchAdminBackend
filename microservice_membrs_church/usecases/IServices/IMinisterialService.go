package IServices

import "github.com/gofiber/fiber/v2"

type IMinisterialService interface {
	GetMinisterialFindAll(c *fiber.Ctx) error
	GetMinisterialFindById(c *fiber.Ctx) error
	CreateMinisterial(c *fiber.Ctx) error
	UpdateMinisterial(c *fiber.Ctx) error
	DeleteMinisterial(c *fiber.Ctx) error
}
