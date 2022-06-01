package internal

import (
	"contactsd/pkg"

	"github.com/gofiber/fiber/v2"
)

type ContactsController struct {
	RepositoryContacts *ContactRepository
}

func NewContactsController(repository *ContactRepository) *ContactsController {
	return &ContactsController{
		RepositoryContacts: repository,
	}
}

func (c *ContactsController) Index(ctx *fiber.Ctx) error {
	results := c.RepositoryContacts.GetAll()
	response := pkg.NewResponseOk(results)
	return ctx.JSON(response)
}

func (c *ContactsController) Detail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result := c.RepositoryContacts.GetById(id)
	response := pkg.NewResponseOk(result)
	return ctx.JSON(response)
}
