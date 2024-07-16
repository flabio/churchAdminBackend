package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"member_church.com/infrastructure/routers"
)

func main() {
	app := fiber.New()
	//app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/swagger/*", swagger.HandlerDefault)
	routers.NewRouter(app)
	routers.NewMinisterialRouter(app)
	routers.NewTeamPescaRouter(app)
	routers.NewChurchRouter(app)
	routers.NewUserRouter(app)
	app.Listen(":8083")
}
