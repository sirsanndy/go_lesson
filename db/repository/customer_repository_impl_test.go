package repository

import (
	"context"
	"db"
	"db/entity"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	customerRepository := NewCustomerRepository(db.GetConnection())
	ctx := context.Background()
	customer := entity.Customer{
		Email:    "test.user123@yahoo.com",
		Name:     "Test Repository",
		UserName: "test.user123",
		IsActive: true,
		Married:  false,
		Balance:  100000,
	}

	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestSelect(t *testing.T) {
	customerRepository := NewCustomerRepository(db.GetConnection())
	ctx := context.Background()

	result, err := customerRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	customerRepository := NewCustomerRepository(db.GetConnection())
	ctx := context.Background()
	result, err := customerRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, customer := range result {
		fmt.Println(customer.Name, customer.UserName)
	}
}
