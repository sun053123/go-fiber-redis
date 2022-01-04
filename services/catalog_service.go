package services

import (
	"github.com/sun053123/go-fiber-redis/entities"
)

type catalogService struct {
	productEnt entities.ProductEntity
}

func NewCatalogService(productEnt entities.ProductEntity) CatalogService {
	return catalogService{productEnt}
}

func (serv catalogService) GetProducts() (products []Product, err error) {

	productsDB, err := serv.productEnt.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product{
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}

	return products, nil
}
