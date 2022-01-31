package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/storyofhis/microservice-go/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		if c.Params("value") != "" {
			return c.SendString("Hello, World 👋!")
		}else {
			return handler.NewErrorMessage(c)
		}
	})
}

func main() {
    app := fiber.New(
		fiber.Config{
			Prefork:true,
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  "Fiber",
			AppName: "Test App v1.0.1",
		},
	)
	
	SetupRoutes(app)

	// Parameters
	app.Get("/:value", handler.GetParams)
	app.Get("/api/list", handler.GetAPI)
	app.Post("/api/register", handler.PostAPI)

    app.Listen(":8080")
}