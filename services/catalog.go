package services

type Product struct {
	Name     string
	Quantity int
}

type CatalogService interface {
	GetProducts() ([]Product, error)
}
