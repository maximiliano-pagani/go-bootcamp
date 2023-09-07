package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const dbPath = "./products.json"

type Product struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price"`
}

var ProductsDB []Product

func InitDB() {
	file, err := os.Open(dbPath)

	if err != nil {
		panic(err)
	}

	json.NewDecoder(file).Decode(&ProductsDB)

	if err != nil {
		panic(err)
	}

	fmt.Println(ProductsDB[0])
}

func GetAll() []Product {
	return ProductsDB
}

func GetById(id int) (Product, error) {
	for _, product := range ProductsDB {
		if product.Id == id {
			return product, nil
		}
	}

	return Product{}, errors.New("Invalid Id")
}

func GetByMinPrice(minPrice float64) []Product {
	var results []Product

	for _, product := range ProductsDB {
		if product.Price > minPrice {
			results = append(results, product)
		}
	}

	return results
}
