package product

import (
	"04-TM-env/internal/domain"
	"encoding/json"
	"errors"
	"os"
)

type ProductDBMemory struct {
	products *[]domain.Product
	lastId   int
}

type ProductRepositoryMemory struct {
	db *ProductDBMemory
}

func NewProductRepositoryMemory(filePath string) *ProductRepositoryMemory {
	database := loadFileData(filePath)
	repository := &ProductRepositoryMemory{db: database}
	return repository
}

func loadFileData(filePath string) *ProductDBMemory {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	db := ProductDBMemory{products: &[]domain.Product{}}
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

func (r *ProductRepositoryMemory) GetPosById(id int) (int, error) {
	var (
		found    bool
		foundPos int
	)

	for pos, product := range *r.db.products {
		if product.Id == id {
			found = true
			foundPos = pos
			break
		}
	}

	if !found {
		return -1, errors.New("Invalid Id")
	}

	return foundPos, nil
}

func (r *ProductRepositoryMemory) GetById(id int) (domain.Product, error) {
	pos, err := r.GetPosById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return (*r.db.products)[pos], nil
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

func (r *ProductRepositoryMemory) Update(product *domain.Product) (domain.Product, error) {
	pos, err := r.GetPosById(product.Id)

	if err != nil {
		return domain.Product{}, err
	}

	(*r.db.products)[pos] = *product
	return *product, nil
}

func (r *ProductRepositoryMemory) Delete(id int) error {
	pos, err := r.GetPosById(id)

	if err != nil {
		return err
	}

	*r.db.products = append((*r.db.products)[:pos], (*r.db.products)[pos+1:]...)
	return nil
}
