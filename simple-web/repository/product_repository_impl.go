package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-web/model"
)

type productRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (repo *productRepositoryImpl) Insert(ctx context.Context, product model.Product) error {
	script := "INSERT INTO (product_name, category, stock, price) VALUES(?, ?, ?, ?)"
	_, err := repo.DB.ExecContext(ctx, script, product.ProductName, product.Category, product.Stock, product.Price)
	if err != nil {
		return err
	}
	return nil
}
func (repo *productRepositoryImpl) Updated(ctx context.Context, product model.Product) error {
	script := "UPDATE t_product SET product_name = ?, category = ?, stock = ?, price = ?, updated_at = ? WHERE id = ?"
	_, err := repo.DB.ExecContext(ctx, script, product.ProductName, product.Category, product.Stock, product.Price, product.UpdatedAt, product.Id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *productRepositoryImpl) Deleted(ctx context.Context, product model.Product) error {
	script := "DELETE FROM t_product WHERE id = ?"
	_, err := repo.DB.ExecContext(ctx, script, product.Id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *productRepositoryImpl) FindById(ctx context.Context, id int32) (model.Product, error) {
	script := "SELECT id, product_name, category, stock, price, create_at, updated_at FROM t_product WHERE id = ? LIMIT 1"
	prod := model.Product{}
	rows, err := repo.DB.QueryContext(ctx, script, id)
	if err != nil {
		return prod, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&prod.Id, &prod.ProductName, &prod.Category, &prod.Stock, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
	} else {
		return prod, errors.New("Product not found")
	}

	return prod, nil
}
func (repo *productRepositoryImpl) FindAll(ctx context.Context) (model.Products, error) {
	script := "SELECT id, product_name, category, stock, price, create_at, updated_at FROM t_product"
	products := model.Products{}
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var prod model.Product
		rows.Scan(&prod.Id, &prod.ProductName, &prod.Category, &prod.Stock, &prod.Price, &prod.CreatedAt, &prod.UpdatedAt)
		products = append(products, prod)
	}

	return products, nil
}
