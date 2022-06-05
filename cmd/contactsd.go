package main

import (
	"contactsd/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* base configuration */
	controllerContacts := internal.NewContactsController()
	/* controllers and routers */
	server := fiber.New(fiber.Config{
		Prefork: true,
	})
	defer server.Listen(":3000")
	v1 := server.Group("/v1")
	v1.Get("/", controllerContacts.Index)
	v1.Get("/:id", controllerContacts.Detail)
}
