package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/storyofhis/microservice-go/database"
	"github.com/storyofhis/microservice-go/handler"
	"github.com/storyofhis/microservice-go/routes"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		if c.Params("value") != "" {
			return c.SendString("Hello, World 👋!")
		}else {
			return handler.NewErrorMessage(c)
		}
	})
	
	// Parameters
	app.Get("/:value", handler.GetParams)
	app.Get("/api/list", handler.GetAPI)
	app.Post("/api/register", handler.PostAPI)

	// user endpoints 
	app.Get("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()

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

    app.Listen(":8080")
}