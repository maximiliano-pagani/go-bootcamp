package mock

import (
	"04-TT-functional-testing/internal/domain"
	"errors"
)

type ProductDBMock struct {
	Products []domain.Product
	Error    error
	LastId   int
}

type ProductRepositoryMock struct {
	db *ProductDBMock
}

func NewProductRepositoryMock(dbMock ProductDBMock) *ProductRepositoryMock {
	repository := &ProductRepositoryMock{db: &dbMock}
	return repository
}

func (r *ProductRepositoryMock) GetAll() ([]domain.Product, error) {
	if r.db.Error != nil {
		return nil, r.db.Error
	}

	return r.db.Products, nil
}

func (r *ProductRepositoryMock) GetPosById(id int) (int, error) {
	var (
		found    bool
		foundPos int
	)

	for pos, product := range r.db.Products {
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

func (r *ProductRepositoryMock) GetById(id int) (domain.Product, error) {
	if r.db.Error != nil {
		return domain.Product{}, r.db.Error
	}

	pos, err := r.GetPosById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return r.db.Products[pos], nil
}

func (r *ProductRepositoryMock) GetByCode(code string) (domain.Product, error) {
	if r.db.Error != nil {
		return domain.Product{}, r.db.Error
	}

	for _, product := range r.db.Products {
		if product.Code == code {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Invalid code")
}

func (r *ProductRepositoryMock) GetByMinPrice(minPrice float64) ([]domain.Product, error) {
	if r.db.Error != nil {
		return nil, r.db.Error
	}

	var results []domain.Product

	for _, product := range r.db.Products {
		if product.Price > minPrice {
			results = append(results, product)
		}
	}

	return results, nil
}

func (r *ProductRepositoryMock) AddNew(product *domain.Product) (domain.Product, error) {
	if r.db.Error != nil {
		return domain.Product{}, r.db.Error
	}

	r.db.LastId++
	product.Id = r.db.LastId
	r.db.Products = append(r.db.Products, *product)
	return *product, nil
}

func (r *ProductRepositoryMock) Update(product *domain.Product) (domain.Product, error) {
	if r.db.Error != nil {
		return domain.Product{}, r.db.Error
	}

	pos, err := r.GetPosById(product.Id)

	if err != nil {
		return domain.Product{}, err
	}

	r.db.Products[pos] = *product
	return *product, nil
}

func (r *ProductRepositoryMock) Delete(id int) error {
	if r.db.Error != nil {
		return r.db.Error
	}

	pos, err := r.GetPosById(id)

	if err != nil {
		return err
	}

	r.db.Products = append(r.db.Products[:pos], r.db.Products[pos+1:]...)
	return nil
}
