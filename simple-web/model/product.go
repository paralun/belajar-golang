package model

import "time"

type Product struct {
	Id          int32
	ProductName string
	Category    string
	Stock       int32
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Products []Product
