package repository

import (
	"context"
	"database/sql"

	"github.com/martinjirku/zasobar/adapters/repository/client"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/pkg/sqlnull"
)

type CategoryRepository struct {
	ctx     context.Context
	queries *client.Queries
}

// we have recursive structure here,
// TODO: refactor db model to handle trees properly https://www.mysqltutorial.org/mysql-adjacency-list-tree/

func NewCategoryRepository(ctx context.Context, db *sql.DB) *CategoryRepository {
	queries := client.New(db)
	return &CategoryRepository{ctx, queries}
}

func (cr *CategoryRepository) List() (entity.Categories, error) {
	categories := []entity.Category{}
	results, err := cr.queries.ListCategories(cr.ctx)
	if err != nil {
		return categories, err
	}
	for _, category := range results {
		categories = append(categories, entity.Category{
			ID:          category.ID,
			Title:       category.Title.String,
			DefaultUnit: category.DefaultUnit.String,
			Path:        entity.CategoryPath(category.Path.String),
		})
	}
	return categories, nil
}

func (cr *CategoryRepository) Create(c entity.Category) (entity.Category, error) {
	ID, err := cr.queries.CreateCategory(cr.ctx, &client.CreateCategoryParams{})
	if err != nil {
		return c, err
	}
	c.ID = int32(ID)
	return c, nil
}

func (cr *CategoryRepository) Update(c entity.Category) (entity.Category, error) {
	err := cr.queries.UpdateCategory(cr.ctx, &client.UpdateCategoryParams{
		ID:          c.ID,
		Title:       sqlnull.FromString(c.Title),
		Path:        sqlnull.FromString(string(c.Path)),
		DefaultUnit: sqlnull.FromString(c.DefaultUnit),
	})
	if err != nil {
		return c, err
	}
	return c, nil
}

func (cr *CategoryRepository) Delete(id int32) error {
	err := cr.queries.DeleteCategory(cr.ctx, id)
	if err != nil {
		return err
	}
	return nil
}
