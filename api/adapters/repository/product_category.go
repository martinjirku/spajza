package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/martinjirku/zasobar/entity"
)

type ProductCategoryRepository struct {
	ctx context.Context
	db  *sql.DB
}

func (pc *ProductCategoryRepository) CreateCategories(categories []entity.ProductCategory) error {
	queryBuilder := strings.Builder{}
	_, err := queryBuilder.WriteString("INSERT INTO product_categories (category_id, name, path, parent_id) VALUES ")
	if err != nil {
		return err
	}
	parameters := make([]interface{}, len(categories)*4)
	for i, category := range categories {
		queryBuilder.WriteString("(?,?,?,?)")
		parameters[i*4+0] = category.CategoryId
		parameters[i*4+1] = category.Name
		parameters[i*4+2] = category.Path
		parameters[i*4+3] = category.ParentId
	}

	pc.db.ExecContext(pc.ctx, queryBuilder.String(), parameters...)
	return nil
}
