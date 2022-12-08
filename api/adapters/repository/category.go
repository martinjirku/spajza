package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/martinjirku/zasobar/entity"
)

type CategoryRepository struct {
	ctx context.Context
	db  *sql.DB
}

// we have recursive structure here,
// TODO: refactor db model to handle trees properly https://www.mysqltutorial.org/mysql-adjacency-list-tree/

func NewCategoryRepository(ctx context.Context, db *sql.DB) *CategoryRepository {
	return &CategoryRepository{ctx, db}
}

func (cr *CategoryRepository) List() (entity.Categories, error) {
	categories := []entity.Category{}

	rows, err := cr.db.QueryContext(cr.ctx, "SELECT id, title, default_unit, path FROM categories WHERE deleted_at IS null")
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c entity.Category
		err := rows.Scan(&c.ID, &c.Title, &c.DefaultUnit, &c.Path)
		if err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (cr *CategoryRepository) Create(c entity.Category) (entity.Category, error) {
	res, err := cr.db.ExecContext(cr.ctx,
		"INSERT INTO categories(created_at, updated_at, title, path, default_unit) VALUES (?,?,?,?,?)",
		time.Now(), time.Now(), c.Title, c.Path, c.DefaultUnit)
	if err != nil {
		return c, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return c, err
	}
	c.ID = id
	return c, nil
}

func (cr *CategoryRepository) Update(c entity.Category) (entity.Category, error) {
	res, err := cr.db.ExecContext(cr.ctx, "UPDATE categories SET updated_at=?,title=?,path=?,default_unit=? WHERE id=?", time.Now(), c.Title, c.Path, c.DefaultUnit, c.ID)
	if err != nil {
		return c, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return c, err
	}
	if affected != 1 {
		return c, entity.ErrNothingUpdated
	}
	return c, nil

}

func (cr *CategoryRepository) Delete(id uint) error {
	res, err := cr.db.ExecContext(cr.ctx, "UPDATE categories SET deleted_at=? WHERE id=?", time.Now(), id)
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
