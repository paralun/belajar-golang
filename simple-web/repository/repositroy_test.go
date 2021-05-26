package repository

import (
	"context"
	"log"
	"simple-web/db"
	"testing"
)

func TestFindAll(t *testing.T) {
	conn, _ := db.GetConnection()
	repo := NewProductRepository(conn)
	ctx := context.Background()
	products, err := repo.FindAll(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println(products)

}
