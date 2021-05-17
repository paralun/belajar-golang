package repository

import (
	"context"
	"go-database/model"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer model.Customer) error
	Update(ctx context.Context, customer model.Customer, id int32) error
	Delete(ctx context.Context, id int32) error
	FindById(ctx context.Context, id int32) (model.Customer, error)
	FindAll(ctx context.Context) ([]model.Customer, error)
}
