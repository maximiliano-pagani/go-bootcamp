package product

import (
	"04-TM-env/internal/domain"
	"encoding/json"
	"errors"
	"os"
)

type ProductDBJson struct {
	jsonPath string
	lastId   int
}

type ProductRepositoryJson struct {
	db *ProductDBJson
}

func NewProductRepositoryJson(jsonPath string) *ProductRepositoryJson {
	database := initJson(jsonPath)
	repository := &ProductRepositoryJson{db: database}
	return repository
}

func initJson(jsonPath string) *ProductDBJson {
	file, err := os.OpenFile(jsonPath, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	db := ProductDBJson{jsonPath: jsonPath}
	products := []domain.Product{}
	json.NewDecoder(file).Decode(&products)

	if err != nil {
		panic(err)
	}

	if len(products) > 0 {
		db.lastId = (products)[len(products)-1].Id
	}

	return &db
}

func (r *ProductRepositoryJson) GetAll() ([]domain.Product, error) {
	file, err := os.Open(r.db.jsonPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	products := []domain.Product{}
	json.NewDecoder(file).Decode(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryJson) GetPosById(products *[]domain.Product, id int) (int, error) {
	var (
		found    bool
		foundPos int
	)

	for pos, product := range *products {
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

func (r *ProductRepositoryJson) GetById(id int) (domain.Product, error) {
	products, err := r.GetAll()

	if err != nil {
		return domain.Product{}, err
	}

	pos, err := r.GetPosById(&products, id)

	if err != nil {
		return domain.Product{}, err
	}

	return products[pos], nil
}

func (r *ProductRepositoryJson) GetByCode(code string) (domain.Product, error) {
	products, err := r.GetAll()

	if err != nil {
		return domain.Product{}, err
	}

	for _, product := range products {
		if product.Code == code {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Invalid code")
}

func (r *ProductRepositoryJson) GetByMinPrice(minPrice float64) ([]domain.Product, error) {
	products, err := r.GetAll()

	if err != nil {
		return nil, err
	}

	var results []domain.Product

	for _, product := range products {
		if product.Price > minPrice {
			results = append(results, product)
		}
	}

	return results, nil
}

func (r *ProductRepositoryJson) AddNew(product *domain.Product) (domain.Product, error) {
	var (
		products []domain.Product
		file     *os.File
		err      error
	)

	if products, err = r.GetAll(); err != nil {
		return domain.Product{}, err
	}

	if file, err = os.Create(r.db.jsonPath); err != nil {
		return domain.Product{}, err
	}

	defer file.Close()

	r.db.lastId++
	product.Id = r.db.lastId
	products = append(products, *product)

	if err = json.NewEncoder(file).Encode(&products); err != nil {
		r.db.lastId--
		return domain.Product{}, err
	}

	return *product, nil
}

func (r *ProductRepositoryJson) Update(product *domain.Product) (domain.Product, error) {
	var (
		products []domain.Product
		pos      int
		file     *os.File
		err      error
	)

	if products, err = r.GetAll(); err != nil {
		return domain.Product{}, err
	}

	if pos, err = r.GetPosById(&products, product.Id); err != nil {
		return domain.Product{}, err
	}

	if file, err = os.Create(r.db.jsonPath); err != nil {
		return domain.Product{}, err
	}

	defer file.Close()

	products[pos] = *product

	if err = json.NewEncoder(file).Encode(&products); err != nil {
		return domain.Product{}, err
	}

	return *product, nil
}

func (r *ProductRepositoryJson) Delete(id int) error {
	var (
		products []domain.Product
		pos      int
		file     *os.File
		err      error
	)

	if products, err = r.GetAll(); err != nil {
		return err
	}

	if pos, err = r.GetPosById(&products, id); err != nil {
		return err
	}

	if file, err = os.Create(r.db.jsonPath); err != nil {
		return err
	}

	defer file.Close()

	products = append(products[:pos], products[pos+1:]...)

	if err = json.NewEncoder(file).Encode(&products); err != nil {
		return err
	}

	return nil
}
