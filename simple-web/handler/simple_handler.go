package handler

import (
	"html/template"
	"net/http"
)

var loadTemplates = template.Must(template.ParseGlob("./templates/*.gohtml"))

func SimpleHandler(writer http.ResponseWriter, request *http.Request) {
	loadTemplates.ExecuteTemplate(writer, "index.gohtml", nil)
}
