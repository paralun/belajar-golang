package go_web

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/upload.gohtml"))
	t.ExecuteTemplate(w, "upload.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request)  {
	//r.ParseMultipartForm(50 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	t := template.Must(template.ParseFiles("./templates/upload_success.gohtml"))
	t.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name" : name,
		"File" : fileHeader.Filename,
	})
}

func TestServerUload(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)

	server := http.Server{
		Addr: "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
