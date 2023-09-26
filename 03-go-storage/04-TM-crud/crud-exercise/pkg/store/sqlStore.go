package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

// NewSqlStore crea un nuevo store de products en una base de datos SQL
func NewSqlStore(storageDB *sql.DB) StoreInterface {
	return &sqlStore{
		db: storageDB,
	}
}

func (s *sqlStore) Read(id int) (domain.Product, error) {
	var product domain.Product

	row := s.db.QueryRow(`SELECT id, name, quantity, code_value, is_published, expiration, price 
						  FROM products WHERE id = ?`, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue,
		&product.IsPublished, &product.Expiration, &product.Price)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *sqlStore) Create(product domain.Product) error {
	stmt, err := s.db.Prepare(`INSERT INTO products (name, quantity, code_value, is_published, expiration, price)
 				  VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	expirationDate, err := time.Parse("02/01/2006", product.Expiration)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue,
		product.IsPublished, expirationDate, product.Price)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	product.Id = int(id)
	return nil
}

func (s *sqlStore) Update(product domain.Product) error {
	stmt, err := s.db.Prepare(`UPDATE products SET name=?, quantity=?, code_value=?,
								is_published=?, expiration=?, price=? WHERE id=?`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	expirationDate, err := time.Parse("02/01/2006", product.Expiration)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.Name, product.Quantity, product.CodeValue,
		product.IsPublished, expirationDate, product.Price, product.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (s *sqlStore) Exists(codeValue string) bool {
	var id int
	row := s.db.QueryRow(`SELECT id FROM products WHERE code_value = ?`, codeValue)
	return !errors.Is(row.Scan(&id), sql.ErrNoRows)
}
