package main

import (
	"contactsd/internal"
	"context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	/* base configuration */
	database := internal.NewMongoDatabase("contacts", context.TODO())
	datasource := internal.NewMongoDataSource(database)
	repositoryContacts := internal.NewContactRepositoryMongo(datasource)
	controllerContacts := internal.NewContactsController(&repositoryContacts)
	/* controllers and routers */
	server := fiber.New(fiber.Config{
		Prefork: false,
	})
	defer server.Listen(":3000")
	v1 := server.Group("/v1")
	v1.Get("/", controllerContacts.Index)
	v1.Get("/:id", controllerContacts.Detail)
}
