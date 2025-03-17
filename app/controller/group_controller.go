package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nastyazz/go_rest_api.git/app/model"
)

func PostGroup(c *fiber.Ctx) error {
	return c.JSON(model.Group{})
}

func GetGroup(c *fiber.Ctx) error {
	return c.JSON([]model.Group{})
}

func PutGroup(c *fiber.Ctx) error {
	return c.JSON(model.Group{})

}

func DeleteGroup(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
