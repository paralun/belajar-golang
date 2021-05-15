package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-test/entity"
	"go-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var catgoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T)  {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := catgoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T)  {
	category := entity.Category{
		Id:   "2",
		Name: "Minuman",
	}
	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := catgoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
}
