package repository

import (
	"encoding/json"
	"errors"
	"os"
)

const dbPath = "./products.json"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductsDB struct {
	products []Product
	lastId   int
}

var db ProductsDB

func InitDB() {
	file, err := os.Open(dbPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	json.NewDecoder(file).Decode(&db.products)

	if err != nil {
		panic(err)
	}

	if len(db.products) > 0 {
		db.lastId = db.products[len(db.products)-1].Id
	}
}

func GetAll() []Product {
	return db.products
}

func GetById(id int) (Product, error) {
	for _, product := range db.products {
		if product.Id == id {
			return product, nil
		}
	}

	return Product{}, errors.New("Invalid Id")
}

func GetByCode(code string) (Product, error) {
	for _, product := range db.products {
		if product.Code == code {
			return product, nil
		}
	}

	return Product{}, errors.New("Invalid code")
}

func GetByMinPrice(minPrice float64) []Product {
	var results []Product

	for _, product := range db.products {
		if product.Price > minPrice {
			results = append(results, product)
		}
	}

	return results
}

func NewProduct(product *Product) (Product, error) {
	db.lastId++
	product.Id = db.lastId
	db.products = append(db.products, *product)
	return *product, nil
}
