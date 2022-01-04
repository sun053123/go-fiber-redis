package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sun053123/go-fiber-redis/services"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv: catalogSrv}
}

func (handl catalogHandler) GetProducts(c *fiber.Ctx) error {

	products, err := handl.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":  "ok",
		"product": products,
	}

	return c.JSON(response)

}
