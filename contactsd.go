package main

import (
	"contactsd/internal"
	"contactsd/pkg"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	injection := internal.GetInjection("contacts", context.TODO())
	/* load repositories */
	repositoryContacts := injection.ContactRepository
	/* configure application/routers */
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	defer app.Listen(":3000")
	app.Get("/", func(c *fiber.Ctx) error {
		results := repositoryContacts.GetAll()
		response := pkg.NewResponseOk(results)
		return c.JSON(response)
	})
	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result := repositoryContacts.GetById(id)
		response := pkg.NewResponseOk(result)
		return c.JSON(response)
	})
}
