package repository

import (
	"context"
	"simple-web/model"
)

type ProductRepository interface {
	Insert(ctx context.Context, product model.Product) error
	Updated(ctx context.Context, product model.Product) error
	Deleted(ctx context.Context, product model.Product) error
	FindById(ctx context.Context, id int32) (model.Product, error)
	FindAll(ctx context.Context) (model.Products, error)
}
