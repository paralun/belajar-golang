package main

import (
	"github.com/gorilla/mux"
	"go-restfull/handler"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/articles", handler.FindAllArticleHandler).Methods("GET")
	r.HandleFunc("/api/articles/{id}", handler.FindByIdArticleHandler).Methods("GET")
	r.HandleFunc("/api/articles", handler.NewArticleHandler).Methods("POST")
	r.HandleFunc("/api/articles/{id}", handler.UpdateArticleHandler).Methods("PUT")
	r.HandleFunc("/api/articles/{id}", handler.DeleteArticleHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7111", r))
}
