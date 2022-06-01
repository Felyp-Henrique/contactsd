package main

import (
	"contactsd/internal"
	"contactsd/pkg"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* base configuration */
	database := internal.NewMongoDatabase("contacts", context.TODO())
	datasource := internal.NewMongoDataSource(database)
	repositoryContacts := internal.NewContactRepositoryMongo(datasource)
	/* controllers and routers */
	server := fiber.New(fiber.Config{
		Prefork: false,
	})
	defer server.Listen(":3000")
	server.Get("/", func(c *fiber.Ctx) error {
		results := repositoryContacts.GetAll()
		response := pkg.NewResponseOk(results)
		return c.JSON(response)
	})
	server.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result := repositoryContacts.GetById(id)
		response := pkg.NewResponseOk(result)
		return c.JSON(response)
	})
}
