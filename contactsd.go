package main

import (
	"contactsd/internal"
	"contactsd/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* configure dependency injection */
	injection := &pkg.Injection{}
	internal.InjectConfigurations(injection)
	internal.InjectionExternal(injection)
	internal.InjectDataSources(injection)
	internal.InjectRepositories(injection)
	/* load repositories */
	repository := injection.Get("repository").(*internal.ContactRepository)
	/* configure application/routers */
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	defer app.Listen(":3000")
	app.Get("/", func(c *fiber.Ctx) error {
		results := repository.GetAll()
		response := pkg.NewResponseOk(results)
		return c.JSON(response)
	})
	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result := repository.GetById(id)
		response := pkg.NewResponseOk(result)
		return c.JSON(response)
	})
}
