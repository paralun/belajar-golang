package service

import (
	"errors"
	"go-restfull/model"
)

var articles model.Articles

type articleServiceImpl struct {
}

func NewArticleService() ArticleService {
	return &articleServiceImpl{}
}

func (svc *articleServiceImpl) New(article model.Article) {
	articles = append(articles, article)
}

func (svc *articleServiceImpl) Update(article model.Article, id string) error {
	for index, item := range articles{
		if item.Id == id {
			//delete
			articles = append(articles[:index], articles[index+1:]...)
			//update
			articles = append(articles, article)
		}
	}
	return nil
}

func (svc *articleServiceImpl) Delete(id string) error {
	for index, item := range articles{
		if item.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
	return nil
}

func (svc *articleServiceImpl) FindById(id string) (model.Article, error) {
	var article = model.Article{}
	for _, item := range articles{
		if item.Id == id {
			article = item
			return article, nil
		}
	}
	return article, errors.New("Not Found")
}

func (svc *articleServiceImpl) FindAll() model.Articles {
	return articles
}
