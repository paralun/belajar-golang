package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Halaman di Redirect ke sini")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request)  {
	//logic
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "https://www.youtube.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr: "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}