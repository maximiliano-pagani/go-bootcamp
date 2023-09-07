package service

import "01-TT-gin/internal/repository"

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

func GetProductsByMinPrice(minPrice float64) []repository.Product {
	return repository.GetByMinPrice(minPrice)
}
