package handler

import "github.com/gofiber/fiber/v2"

func NewErrorMessage (c *fiber.Ctx) error {
	if c.Params("value") != "" {
		return c.SendString("Hello world")
	}
	return fiber.NewError(782, "Custom error message")
}

func GetAPI (c *fiber.Ctx) error {
	return c.SendString("I'm a GET Request")
}

func GetParams(c *fiber.Ctx) error {
	value := c.Params("value")
	if value != "" {
		return c.SendString("Hello i'm " + value)
	}
	return c.SendString("doesn't getting the value")
}

func PostAPI(c *fiber.Ctx) error {
	return c.SendString("I'm a POST request")
}