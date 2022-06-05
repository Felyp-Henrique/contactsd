package main

import (
	"contactsd/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* controllers */
	controllerContacts := internal.NewContactsController()
	/* routers */
	configuration := fiber.Config{
		Prefork: true,
	}
	server := fiber.New(configuration)
	defer server.Listen(":3000")
	v1 := server.Group("/v1")
	v1.Get("/", controllerContacts.Index)
	v1.Get("/:id", controllerContacts.Detail)
	v1.Post("/", controllerContacts.Create)
	v1.Put("/", controllerContacts.Update)
}
