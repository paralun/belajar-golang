package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

var bookstore = make(map[string]*Book)

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprint(writer, "Welcome!\n")
}

func NewBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	book := &Book{}
	if err := populateModel(writer, request, params, book); err != nil {
		writeErrorResponse(writer, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	bookstore[book.ISDN] = book
	writeOKResponse(writer, book)
}

func FindAllBook(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	var books []*Book
	for _, book := range bookstore {
		books = append(books, book)
	}
	writeOKResponse(writer, books)
}

func FindBookById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	isdn := params.ByName("isdn")
	book, ok := bookstore[isdn]
	if !ok {
		writeErrorResponse(writer, http.StatusNotFound, "Record Not Found")
		return
	}
	writeOKResponse(writer, book)
}

func populateModel(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JsonResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}


