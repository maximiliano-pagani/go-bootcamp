package product

import "04-TM-env/internal/domain"

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	GetByCode(code string) (domain.Product, error)
	GetByMinPrice(minPrice float64) ([]domain.Product, error)
	AddNew(product *domain.Product) (domain.Product, error)
	Update(product *domain.Product) (domain.Product, error)
	Delete(id int) error
}
