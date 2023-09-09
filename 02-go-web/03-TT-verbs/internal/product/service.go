package product

import "03-TT-verbs/internal/domain"

type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductByCode(code string) (domain.Product, error)
	GetProductsByMinPrice(minPrice float64) ([]domain.Product, error)
	NewProduct(product *domain.Product) (domain.Product, error)
	ReplaceProduct(product *domain.Product) (domain.Product, error)
	UpdateProduct(product *domain.Product) (domain.Product, error)
	DeleteProduct(id int) error
}
