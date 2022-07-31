package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinjirku/zasobar/domain"
)

type CategoryRepository struct {
	db *sql.DB
}

const (
	listAllStmt         = "SELECT id, created_at, updated_at, deleted_at, title, default_unit, path FROM categories WHERE deleted_at IS null"
	insertCategory5Stmt = "INSERT INTO categories(created_at, updated_at, title, path, default_unit) VALUES (?,?,?,?,?)"
	updateCategory5Stmt = "UPDATE categories SET updated_at=?,title=?,path=?,default_unit=? WHERE id=?"
	deleteCategory2Stmt = "UPDATE categories SET deleted_at=? WHERE id=?"
)

// we have recursive structure here,
// TODO: refactor db model to handle trees properly https://www.mysqltutorial.org/mysql-adjacency-list-tree/

func NewCategoryService(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (c *CategoryRepository) ListAll(ctx context.Context) ([]domain.Category, error) {
	categories := []domain.Category{}

	rows, err := c.db.QueryContext(ctx, listAllStmt)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.Category
		err := rows.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt, &c.Title, &c.DefaultUnit, &c.Path)
		if err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (cs *CategoryRepository) CreateItem(ctx context.Context, c domain.Category) (domain.Category, error) {
	res, err := cs.db.ExecContext(ctx, insertCategory5Stmt, time.Now(), time.Now(), c.Title, c.Path, c.DefaultUnit)
	if err != nil {
		return c, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return c, err
	}
	c.ID = uint(id)
	return c, nil
}

func (cs *CategoryRepository) UpdateItem(ctx context.Context, c domain.Category) (domain.Category, error) {
	c.UpdatedAt = time.Now()
	res, err := cs.db.ExecContext(ctx, updateCategory5Stmt, c.UpdatedAt, c.Title, c.Path, c.DefaultUnit, c.ID)
	if err != nil {
		return c, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return c, err
	}
	if affected != 1 {
		return c, errors.New("nothing updated")
	}
	return c, nil

}

func (cs *CategoryRepository) DeleteItem(ctx context.Context, id uint) error {
	res, err := cs.db.ExecContext(ctx, deleteCategory2Stmt, time.Now(), id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("nothing deleted")
	}

	return nil
}
