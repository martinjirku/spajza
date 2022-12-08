package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type CategoryRepository interface {
	List() (entity.Categories, error)
	Create(c entity.Category) (entity.Category, error)
	Update(c entity.Category) (entity.Category, error)
	Delete(id uint) error
}

type CategoryUsecase struct {
	repository CategoryRepository
}

func CreateCategoryUsecase(categoryRepository CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{categoryRepository}
}

func (cs *CategoryUsecase) List() (entity.Categories, error) {
	return cs.repository.List()
}

func (cs *CategoryUsecase) Create(c entity.Category) (entity.Category, error) {
	categories, err := cs.repository.List()
	if err != nil {
		return entity.Category{}, err
	}
	err = categories.AddCategory(c)
	if err != nil {
		return entity.Category{}, err
	}
	return cs.repository.Create(c)
}

func (cs *CategoryUsecase) Update(c entity.Category) (entity.Category, error) {
	if !c.Validate() {
		return c, entity.ErrInvalidEntity
	}
	return cs.repository.Update(c)
}

func (cs *CategoryUsecase) Delete(id uint) error {
	return cs.repository.Delete(id)
}
