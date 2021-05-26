package handler

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"simple-web/db"
	"simple-web/model"
	"simple-web/repository"
	"strconv"
)

var connection *sql.DB
var templates = template.Must(template.ParseGlob("./templates/*.gohtml"))

func SetupConn() (err error) {
	connection, err = db.GetConnection()
	return
}

func CloseConn() {
	if connection != nil {
		connection.Close()
	}
}

func IndexHandler(writer http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	repo := repository.NewProductRepository(connection)
	products, err := repo.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	templates.ExecuteTemplate(writer, "index.gohtml", products)
}

func AddHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.FormValue("productName")
	category := request.FormValue("category")
	stock := request.FormValue("stock")
	stock32, _ := strconv.ParseInt(stock, 10, 32)
	price := request.FormValue("price")
	price64, _ := strconv.ParseInt(price, 10, 64)

	product := model.Product{
		ProductName: name,
		Category:    category,
		Stock:       int32(stock32),
		Price:       price64,
	}

	ctx := context.Background()
	repo := repository.NewProductRepository(connection)

	err := repo.Insert(ctx, product)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
