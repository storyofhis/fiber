package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

	// Simple route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	// Parameters
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
		  return c.SendString("Hello " + c.Params("name"))
		  // => Hello john
		}
		return c.SendString("Where is john?")
	})
    app.Listen(":8080")
}