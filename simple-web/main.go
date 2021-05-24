package main

import (
	"log"
	"net/http"
	"simple-web/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.SimpleHandler)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:9292",
		Handler: mux,
	}

	log.Println("Server running on port 9292")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
