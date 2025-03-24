package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nastyazz/go_rest_api.git/app/controller"
)

func GroupRoutes(a *fiber.App) {
	routes := a.Group("/api/v1/group")

	routes.Get("/", controller.GetGroup)
	routes.Post("/", controller.PostGroup)
	routes.Put("/:id", controller.PutGroup)
	routes.Delete("/:id", controller.DeleteGroup)

}
