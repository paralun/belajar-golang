package repository

import (
	"context"
	"fmt"
	"go-database/model"
	"go-database/mysql"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db := mysql.GetConnectionMysql()
	repo := NewCustomerRepository(db)
	defer db.Close()

	ctx := context.Background()

	layout := "2006-01-02"
	birth, _ := time.Parse(layout, "1990-06-20")
	customer := model.Customer{
		Name:      "Gunawan",
		Email:     "gun@gmail.com",
		Balance:   8000000,
		Rating:    90.7,
		BirthDate: birth,
		Married:   true,
	}

	err := repo.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert Success")
}

func TestUpdate(t *testing.T) {
	db := mysql.GetConnectionMysql()
	repo := NewCustomerRepository(db)
	defer db.Close()

	ctx := context.Background()

	layout := "2006-01-02"
	birth, _ := time.Parse(layout, "1998-08-12")
	customer := model.Customer{
		Id:        4,
		Name:      "Fitri",
		Email:     "fit@gmail.com",
		Balance:   90334400,
		Rating:    99.5,
		BirthDate: birth,
		Married:   false,
	}

	err := repo.Update(ctx, customer, customer.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Success")
}

func TestDelete(t *testing.T) {
	db := mysql.GetConnectionMysql()
	repo := NewCustomerRepository(db)
	defer db.Close()

	ctx := context.Background()

	err := repo.Delete(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete Success")
}

func TestFindById(t *testing.T) {
	db := mysql.GetConnectionMysql()
	repo := NewCustomerRepository(db)
	defer db.Close()

	ctx := context.Background()

	cust, err := repo.FindById(ctx, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(cust)
}

func TestFindAll(t *testing.T) {
	db := mysql.GetConnectionMysql()
	repo := NewCustomerRepository(db)
	defer db.Close()

	ctx := context.Background()

	custs, err := repo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(custs)
}
