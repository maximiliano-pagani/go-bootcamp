package product

import "03-TM-structure/internal/domain"

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	GetByCode(code string) (domain.Product, error)
	GetByMinPrice(minPrice float64) ([]domain.Product, error)
	AddNew(product *domain.Product) (domain.Product, error)
}
