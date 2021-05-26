package main

import (
	"log"
	"net/http"
	"simple-web/handler"
)

func main() {
	log.Println("Setup database connection")
	err := handler.SetupConn()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.CloseConn()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/add", handler.AddHandler)

	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:9292",
		Handler: mux,
	}

	log.Println("Server running on port 9292")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
