package go_web

import (
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request)  {
	file := "sample.jpeg"
	w.Header().Add("Content-Disposition", "attachment; filename=\"" + file + "\"")
	http.ServeFile(w, r, "./resources/" + file)
}

func TestDownloadFile(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)

	server := http.Server{
		Addr: "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
