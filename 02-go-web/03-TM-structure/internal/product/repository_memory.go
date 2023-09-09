package product

import (
	"03-TM-structure/internal/domain"
	"encoding/json"
	"errors"
	"os"
)

type ProductDB struct {
	products *[]domain.Product
	lastId   int
}

type ProductRepositoryMemory struct {
	db *ProductDB
}

func NewProductRepositoryMemory(filePath string) *ProductRepositoryMemory {
	database := loadFileData(filePath)
	repository := &ProductRepositoryMemory{db: database}
	return repository
}

func loadFileData(filePath string) *ProductDB {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	db := ProductDB{products: &[]domain.Product{}}
	json.NewDecoder(file).Decode(db.products)

	if err != nil {
		panic(err)
	}

	if len(*db.products) > 0 {
		db.lastId = (*db.products)[len(*db.products)-1].Id
	}

	return &db
}

func (r *ProductRepositoryMemory) GetAll() ([]domain.Product, error) {
	return *r.db.products, nil
}

func (r *ProductRepositoryMemory) GetById(id int) (domain.Product, error) {
	for _, product := range *r.db.products {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Invalid Id")
}

func (r *ProductRepositoryMemory) GetByCode(code string) (domain.Product, error) {
	for _, product := range *r.db.products {
		if product.Code == code {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Invalid code")
}

func (r *ProductRepositoryMemory) GetByMinPrice(minPrice float64) ([]domain.Product, error) {
	var results []domain.Product

	for _, product := range *r.db.products {
		if product.Price > minPrice {
			results = append(results, product)
		}
	}

	return results, nil
}

func (r *ProductRepositoryMemory) AddNew(product *domain.Product) (domain.Product, error) {
	r.db.lastId++
	product.Id = r.db.lastId
	*r.db.products = append(*r.db.products, *product)
	return *product, nil
}
