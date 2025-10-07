package repository

import (
	"context"
	"db/entity"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	FindById(ctx context.Context, id int32) (entity.Customer, error)
	FindAll(ctx context.Context) ([]entity.Customer, error)
}
