package internal

import (
	"contactsd/pkg"
	"context"

	"github.com/gofiber/fiber/v2"
)

type ContactsController struct{}

func NewContactsController() *ContactsController {
	return &ContactsController{}
}

func (c *ContactsController) Index(ctx *fiber.Ctx) error {
	if repository, err := c.repository(); err == nil {
		results, _ := repository.GetAll()
		response := pkg.NewResponseOk(results)
		return ctx.JSON(response)
	} else {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
}

func (c *ContactsController) Detail(ctx *fiber.Ctx) error {
	if repository, err := c.repository(); err == nil {
		id := ctx.Params("id")
		result, _ := repository.GetById(id)
		response := pkg.NewResponseOk(result)
		return ctx.JSON(response)
	} else {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
}

func (c *ContactsController) Create(ctx *fiber.Ctx) error {
	contact := pkg.Contact{}
	if err := ctx.BodyParser(&contact); err != nil {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
	if repository, err := c.repository(); err == nil {
		repository.Insert(contact)
		response := pkg.NewResponseOk(contact)
		return ctx.JSON(response)
	} else {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
}

func (c *ContactsController) Update(ctx *fiber.Ctx) error {
	contact := pkg.Contact{}
	if err := ctx.BodyParser(&contact); err != nil {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
	if repository, err := c.repository(); err == nil {
		repository.Update(contact)
		response := pkg.NewResponseOk(contact)
		return ctx.JSON(response)
	} else {
		response := pkg.NewResponseInternalError(err)
		return ctx.JSON(response)
	}
}

func (c *ContactsController) repository() (*ContactsRepository, error) {
	context := context.TODO()
	if database, err := pkg.NewMongo(context, ""); err == nil {
		datasource := NewContactsMongoDataSource(context, database)
		repository := NewContactsRepositoryMongo(datasource)
		return &repository, nil
	} else {
		return nil, err
	}
}
