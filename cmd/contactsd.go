package main

import (
	"contactsd/internal"
	"contactsd/pkg"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* base configuration */
	database := pkg.NewMongoDatabase("contacts", context.TODO())
	datasource := pkg.NewMongoDataSource(database)
	repositoryContacts := internal.NewContactsRepositoryMongo(datasource)
	controllerContacts := internal.NewContactsController(&repositoryContacts)
	/* controllers and routers */
	server := fiber.New(fiber.Config{
		Prefork: true,
	})
	defer server.Listen(":3000")
	v1 := server.Group("/v1")
	v1.Get("/", controllerContacts.Index)
	v1.Get("/:id", controllerContacts.Detail)
}
