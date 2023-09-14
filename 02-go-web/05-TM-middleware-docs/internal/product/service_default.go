package product

import (
	"04-TT-functional-testing/internal/domain"
	"errors"
)

type ProductServiceDefault struct {
	repository ProductRepository
}

func NewProductServiceDefault(repository ProductRepository) *ProductServiceDefault {
	service := &ProductServiceDefault{repository: repository}
	return service
}

func (s *ProductServiceDefault) GetAllProducts() ([]domain.Product, error) {
	return s.repository.GetAll()
}

func (s *ProductServiceDefault) GetProductById(id int) (domain.Product, error) {
	product, err := s.repository.GetById(id)
	return product, err
}

func (s *ProductServiceDefault) GetProductByCode(code string) (domain.Product, error) {
	product, err := s.repository.GetByCode(code)
	return product, err
}

func (s *ProductServiceDefault) GetProductsByMinPrice(minPrice float64) ([]domain.Product, error) {
	return s.repository.GetByMinPrice(minPrice)
}

func (s *ProductServiceDefault) NewProduct(product *domain.Product) (domain.Product, error) {
	if err := s.validateNewProduct(product); err != nil {
		return domain.Product{}, err
	}

	addedProduct, err := s.repository.AddNew(product)

	if err != nil {
		return domain.Product{}, err
	}

	return addedProduct, nil
}

func (s *ProductServiceDefault) validateNewProduct(product *domain.Product) error {
	if _, err := s.GetProductByCode(product.Code); err == nil {
		return errors.New("Invalid product code: already exists")
	}

	return nil
}

func (s *ProductServiceDefault) ReplaceProduct(product *domain.Product) (domain.Product, error) {
	if foundProduct, err := s.GetProductByCode(product.Code); err == nil && foundProduct.Id != product.Id {
		return domain.Product{}, errors.New("Invalid product code: already exists")
	}

	replacedProduct, err := s.repository.Update(product)
	return replacedProduct, err
}

func (s *ProductServiceDefault) UpdateProduct(product *domain.Product) (domain.Product, error) {
	if foundProduct, err := s.GetProductByCode(product.Code); err == nil && foundProduct.Id != product.Id {
		return domain.Product{}, errors.New("Invalid product code: already exists")
	}

	updatedProduct, err := s.repository.Update(product)
	return updatedProduct, err
}

func (s *ProductServiceDefault) DeleteProduct(id int) error {
	err := s.repository.Delete(id)
	return err
}
