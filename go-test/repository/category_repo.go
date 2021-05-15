package repository

import "go-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
