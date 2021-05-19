package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:7000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Belajar Golang Web")
	}

	server := http.Server{
		Addr:    "localhost:7000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}