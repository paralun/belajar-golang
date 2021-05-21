package go_web

import (
	"embed"
	"html/template"
	"net/http"
	"testing"
)

//go:embed templates
var tempEmbed embed.FS
var myTempEmbed = template.Must(template.ParseFS(tempEmbed, "templates/*.gohtml"))

//var myTempFiles = template.Must(template.ParseGlob("./templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request)  {
	myTempEmbed.ExecuteTemplate(w, "simple.gohtml", "Ini template caching")
	//myTempFiles.ExecuteTemplate(w, "simple.gohtml", "Ini template caching")
}

func TestServerTemplateCaching(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/simple-caching", TemplateCaching)

	server := http.Server{
		Addr: "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
