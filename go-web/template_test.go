package go_web

import (
	"embed"
	"html/template"
	"net/http"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request)  {
	templateText := "<html><body>{{.}}</body></html>"
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hallo Template")
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hallo Template File")
}

func SimpleHTMLDir(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hallo Template Directory")
}

//go:embed templates/*.gohtml
var temp embed.FS

func SimpleHTMLEmbed(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFS(temp, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hallo Template Embed")
}

func SimpleDataMap(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "param.gohtml", map[string]interface{}{
		"Title" : "Data Map",
		"Name" : "Kusmambang",
	})
}

type Address struct {
	Street string
}

type Page struct {
	Title string
	Name string
	Address Address
}

func SimpleDataStruct(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "param.gohtml", Page{
		Title: "Data Struct",
		Name:  "Paralun",
		Address:Address{Street: "Jl Mana Aja"},
	})
}

func TestServerTemplate(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/simple", SimpleHTML)
	mux.HandleFunc("/simple-file", SimpleHTMLFile)
	mux.HandleFunc("/simple-dir", SimpleHTMLDir)
	mux.HandleFunc("/simple-embed", SimpleHTMLEmbed)
	mux.HandleFunc("/simple-map", SimpleDataMap)
	mux.HandleFunc("/simple-struct", SimpleDataStruct)

	server := http.Server{
		Addr: "localhost:7000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
