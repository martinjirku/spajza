package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type CategoryRepository interface {
	ListAll() ([]entity.Category, error)
	CreateItem(c entity.Category) (entity.Category, error)
	UpdateItem(c entity.Category) (entity.Category, error)
	DeleteItem(id uint) error
}

type CategoryService struct {
	repository CategoryRepository
}

func CreateCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository}
}

func (cs *CategoryService) ListAll() ([]entity.Category, error) {
	return cs.repository.ListAll()
}

func (cs *CategoryService) CreateItem(c entity.Category) (entity.Category, error) {
	return cs.repository.CreateItem(c)
}

func (cs *CategoryService) UpdateItem(c entity.Category) (entity.Category, error) {
	return cs.repository.UpdateItem(c)
}

func (cs *CategoryService) DeleteItem(id uint) error {
	return cs.repository.DeleteItem(id)
}
