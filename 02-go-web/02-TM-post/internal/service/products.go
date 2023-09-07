package service

import (
	"02-TM-post/internal/repository"
	"errors"
)

func Init() {
	repository.InitDB()
}

func GetAllProducts() []repository.Product {
	return repository.GetAll()
}

func GetProductById(id int) (product repository.Product, err error) {
	product, err = repository.GetById(id)
	return
}

func GetProductByCode(code string) (product repository.Product, err error) {
	product, err = repository.GetByCode(code)
	return
}

func GetProductsByMinPrice(minPrice float64) []repository.Product {
	return repository.GetByMinPrice(minPrice)
}

func NewProduct(product *repository.Product) (repository.Product, error) {
	if err := validateNewProduct(product); err != nil {
		return repository.Product{}, err
	}

	addedProduct, err := repository.NewProduct(product)

	if err != nil {
		return repository.Product{}, err
	}
	
	return addedProduct, nil
}

func validateNewProduct(product *repository.Product) error {
	if _, err := GetProductByCode(product.Code); err == nil {
		return errors.New("Invalid product code: already exists")
	}

	return nil
}
