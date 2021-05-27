package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-restfull/model"
	"go-restfull/service"
	"net/http"
)

var svc = service.NewArticleService()

func NewArticleHandler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")

	var article model.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	if err != nil{
		json.NewEncoder(writer).Encode(err)
		return
	}

	svc.New(article)
	json.NewEncoder(writer).Encode(map[string]string{
		"status" : "created",
	})
}

func UpdateArticleHandler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")

	var article model.Article
	params := mux.Vars(request)
	id := params["id"]

	err := json.NewDecoder(request.Body).Decode(&article)
	if err != nil{
		json.NewEncoder(writer).Encode(err)
		return
	}

	errUpd := svc.Update(article, id)
	if errUpd != nil{
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"status" : "updated",
	})

}

func DeleteArticleHandler(writer http.ResponseWriter, request *http.Request)  {
	params := mux.Vars(request)
	id := params["id"]

	err := svc.Delete(id)
	if err != nil{
		json.NewEncoder(writer).Encode(err)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"status" : "deleted",
	})
}

func FindByIdArticleHandler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	id := params["id"]
	article, err := svc.FindById(id)

	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{
			"error" : "not found",
		})
		return
	}

	json.NewEncoder(writer).Encode(article)
}
func FindAllArticleHandler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")

	articles := svc.FindAll()
	json.NewEncoder(writer).Encode(articles)
}
