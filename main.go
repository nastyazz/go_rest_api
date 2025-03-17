package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nastyazz/go_rest_api.git/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	routes.ContactRoutes(app)
	routes.GroupRoutes(app)

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
