package repository

import (
	"context"
	"database/sql"

	"github.com/martinjirku/zasobar/adapters/repository/client"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/pkg/sqlnull"
)

type ProductCategoryRepository struct {
	ctx     context.Context
	db      *sql.DB
	queries *client.Queries
}

func NewProductCategoryRepository(ctx context.Context, db *sql.DB) *ProductCategoryRepository {
	queries := client.New(db)
	return &ProductCategoryRepository{ctx, db, queries}
}

func (pc *ProductCategoryRepository) CreateCategories(categories []entity.ProductCategory) error {
	tx, err := pc.db.BeginTx(pc.ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := pc.queries.WithTx(tx)
	for _, category := range categories {
		_, err := q.InsertProductCategory(pc.ctx, &client.InsertProductCategoryParams{
			CategoryID: category.CategoryId,
			Name:       sqlnull.FromString(category.Name),
			Path:       sqlnull.FromString(string(category.Path)),
			ParentID:   sqlnull.FromInt32Ptr(category.ParentId),
		})

		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
