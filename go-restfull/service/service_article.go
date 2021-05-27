package service

import "go-restfull/model"

type ArticleService interface {
	New(article model.Article)
	Update(article model.Article, id string) error
	Delete(id string) error
	FindById(id string) (model.Article, error)
	FindAll() model.Articles
}
