package usecases

import (
	"context"

	"github.com/martinjirku/zasobar/entity"
)

type CategoryRepository interface {
	ListAll(ctx context.Context) ([]entity.Category, error)
	CreateItem(ctx context.Context, c entity.Category) (entity.Category, error)
	UpdateItem(ctx context.Context, c entity.Category) (entity.Category, error)
	DeleteItem(ctx context.Context, id uint) error
}

type CategoryService struct {
	repository CategoryRepository
}

func CreateCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository}
}

func (cs *CategoryService) ListAll(ctx context.Context) ([]entity.Category, error) {
	return cs.repository.ListAll(ctx)
}

func (cs *CategoryService) CreateItem(ctx context.Context, c entity.Category) (entity.Category, error) {
	return cs.repository.CreateItem(ctx, c)
}

func (cs *CategoryService) UpdateItem(ctx context.Context, c entity.Category) (entity.Category, error) {
	return cs.repository.UpdateItem(ctx, c)
}

func (cs *CategoryService) DeleteItem(ctx context.Context, id uint) error {
	return cs.repository.DeleteItem(ctx, id)
}
