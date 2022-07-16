package categories

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type CategoryService struct {
	db *sql.DB
}

const (
	listAllStmt = "SELECT id, created_at, updated_at, deleted_at, title, default_unit, path FROM categories WHERE deleted_at IS null"
)

func NewCategoryService(db *sql.DB) CategoryService {
	return CategoryService{db: db}
}

func (cs *CategoryService) ListAll(ctx context.Context) ([]Category, error) {
	categories := []Category{}

	rows, err := cs.db.QueryContext(ctx, listAllStmt)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt, &c.Title, &c.DefaultUnit, &c.Path)
		if err != nil {
			return categories, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (cs *CategoryService) CreateItem(ctx context.Context, c Category) (Category, error) {
	res, err := cs.db.ExecContext(ctx, "INSERT INTO categories(created_at, updated_at, title, path, default_unit) VALUES (?,?,?,?,?)",
		time.Now(), time.Now(), c.Title, c.Path, c.DefaultUnit)
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

func (cs *CategoryService) UpdateItem(ctx context.Context, c Category) (Category, error) {
	c.UpdatedAt = time.Now()
	res, err := cs.db.ExecContext(ctx, "UPDATE categories SET updated_at=?,title=?,path=?,default_unit=? WHERE id=?", c.UpdatedAt, c.Title, c.Path, c.DefaultUnit, c.ID)
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
func (cs *CategoryService) DeleteItem(ctx context.Context, c Category) error {
	res, err := cs.db.ExecContext(ctx, "UPDATE categories SET deleted_at=? WHERE id=?", time.Now(), c.ID)
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
