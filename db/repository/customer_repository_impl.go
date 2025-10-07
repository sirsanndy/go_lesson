package repository

import (
	"context"
	"database/sql"
	"db/entity"
	"errors"
	"strconv"
)

type customerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImpl{DB: db}
}

func (repository *customerRepositoryImpl) Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := "INSERT INTO customer(email, name, user_name, is_active, married, balance) VALUES($1, $2, $3, $4, $5, $6)"
	result, err := repository.DB.ExecContext(ctx, script, customer.Email, customer.Name, customer.UserName, customer.IsActive, customer.Married, customer.Balance)
	if err != nil {
		return customer, err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return customer, err
	}

	customer.Id = int32(id)
	return customer, nil
}

func (repository *customerRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Customer, error) {
	script := "SELECT id, email, name FROM customer WHERE id = $1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return entity.Customer{}, err
	}

	defer rows.Close()
	customer := entity.Customer{}
	if rows.Next() {
		if err := rows.Scan(&customer.Id, &customer.Email, &customer.Name); err != nil {
			return customer, err
		}
		return customer, nil
	} else {
		return customer, errors.New("ID " + strconv.Itoa(int(id)))
	}
}

func (repository *customerRepositoryImpl) FindAll(ctx context.Context) ([]entity.Customer, error) {
	script := "SELECT id, email, name FROM customer"
	rows, err := repository.DB.QueryContext(ctx, script)
	var customers []entity.Customer
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		customer := entity.Customer{}
		if err := rows.Scan(&customer.Id, &customer.Email, &customer.Name); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
