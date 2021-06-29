package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/books", FindAllBook)
	router.GET("/api/books/:isdn", FindBookById)
	router.POST("/api/books", NewBook)

	auth := &AuthMiddleware{
		Handler: router,
		User:    "admin",
		Pass:    "admin123",
	}

	log.Fatal(http.ListenAndServe(":9191", auth))
}
