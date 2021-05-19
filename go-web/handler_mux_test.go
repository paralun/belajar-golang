package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halaman Index")
	})

	mux.HandleFunc("/hallo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hallo, lagi belajar golang web nihh..")
	})

	mux.HandleFunc("/produk/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halaman Produk")
	})

	mux.HandleFunc("/produk/diskon/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halaman Produk Diskon")
	})

	server := http.Server{
		Addr:    "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
