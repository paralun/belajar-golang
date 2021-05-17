package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/model"
)

type customerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository  {
	return &customerRepositoryImpl{DB: db}
}

func (repo *customerRepositoryImpl) Insert(ctx context.Context, customer model.Customer) error {
	script := "INSERT INTO customer (name, email, balance, rating, birth_date, married) VALUES(?,?,?,?,?,?)"
	_, err := repo.DB.ExecContext(ctx, script,
		customer.Name, customer.Email, customer.Balance, customer.Rating, customer.BirthDate, customer.Married)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerRepositoryImpl) Update(ctx context.Context, customer model.Customer, id int32) error {
	script := "UPDATE customer SET name = ?, email = ?, balance = ?, rating = ?, birth_date = ?, married = ? WHERE id = ?"
	_, err := repo.DB.ExecContext(ctx, script,
		customer.Name, customer.Email, customer.Balance, customer.Rating, customer.BirthDate, customer.Married, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerRepositoryImpl) Delete(ctx context.Context, id int32) error {
	script := "DELETE FROM customer WHERE id = ?"
	_, err := repo.DB.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerRepositoryImpl) FindById(ctx context.Context, id int32) (model.Customer, error) {
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer WHERE id = ?"
	cust := model.Customer{}
	rows, err := repo.DB.QueryContext(ctx, script, id)
	if err != nil {
		return cust, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&cust.Id, &cust.Name, &cust.Email, &cust.Balance, &cust.Rating, &cust.BirthDate, &cust.Married, &cust.Married)
	} else {
		return cust, errors.New("Not Found")
	}

	return cust, nil
}

func (repo *customerRepositoryImpl) FindAll(ctx context.Context) ([]model.Customer, error) {
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	var customers []model.Customer
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return customers, err
	}
	defer rows.Close()

	for rows.Next() {
		var cust model.Customer
		rows.Scan(&cust.Id, &cust.Name, &cust.Email, &cust.Balance, &cust.Rating, &cust.BirthDate, &cust.Married, &cust.Married)
		customers = append(customers, cust)
	}

	return customers, nil
}