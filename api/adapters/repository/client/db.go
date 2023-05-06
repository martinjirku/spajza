// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package client

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createProductCategoryStmt, err = db.PrepareContext(ctx, createProductCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProductCategory: %w", err)
	}
	if q.createStorageConsumptionStmt, err = db.PrepareContext(ctx, createStorageConsumption); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStorageConsumption: %w", err)
	}
	if q.createStorageItemStmt, err = db.PrepareContext(ctx, createStorageItem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStorageItem: %w", err)
	}
	if q.createStoragePlaceStmt, err = db.PrepareContext(ctx, createStoragePlace); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStoragePlace: %w", err)
	}
	if q.deleteCategoryStmt, err = db.PrepareContext(ctx, deleteCategory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategory: %w", err)
	}
	if q.deleteStoragePlaceStmt, err = db.PrepareContext(ctx, deleteStoragePlace); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteStoragePlace: %w", err)
	}
	if q.getStorageConsumptionByIdStmt, err = db.PrepareContext(ctx, getStorageConsumptionById); err != nil {
		return nil, fmt.Errorf("error preparing query GetStorageConsumptionById: %w", err)
	}
	if q.getStorageItemByIdStmt, err = db.PrepareContext(ctx, getStorageItemById); err != nil {
		return nil, fmt.Errorf("error preparing query GetStorageItemById: %w", err)
	}
	if q.getStoragePlaceByIdStmt, err = db.PrepareContext(ctx, getStoragePlaceById); err != nil {
		return nil, fmt.Errorf("error preparing query GetStoragePlaceById: %w", err)
	}
	if q.listCategoriesStmt, err = db.PrepareContext(ctx, listCategories); err != nil {
		return nil, fmt.Errorf("error preparing query ListCategories: %w", err)
	}
	if q.listStorageConsumptionsStmt, err = db.PrepareContext(ctx, listStorageConsumptions); err != nil {
		return nil, fmt.Errorf("error preparing query ListStorageConsumptions: %w", err)
	}
	if q.listStorageItemsStmt, err = db.PrepareContext(ctx, listStorageItems); err != nil {
		return nil, fmt.Errorf("error preparing query ListStorageItems: %w", err)
	}
	if q.listStoragePlacesStmt, err = db.PrepareContext(ctx, listStoragePlaces); err != nil {
		return nil, fmt.Errorf("error preparing query ListStoragePlaces: %w", err)
	}
	if q.updateCategoryStmt, err = db.PrepareContext(ctx, updateCategory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategory: %w", err)
	}
	if q.updateStorageItemStmt, err = db.PrepareContext(ctx, updateStorageItem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStorageItem: %w", err)
	}
	if q.updateStoragePlaceStmt, err = db.PrepareContext(ctx, updateStoragePlace); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStoragePlace: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createProductCategoryStmt != nil {
		if cerr := q.createProductCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductCategoryStmt: %w", cerr)
		}
	}
	if q.createStorageConsumptionStmt != nil {
		if cerr := q.createStorageConsumptionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStorageConsumptionStmt: %w", cerr)
		}
	}
	if q.createStorageItemStmt != nil {
		if cerr := q.createStorageItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStorageItemStmt: %w", cerr)
		}
	}
	if q.createStoragePlaceStmt != nil {
		if cerr := q.createStoragePlaceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStoragePlaceStmt: %w", cerr)
		}
	}
	if q.deleteCategoryStmt != nil {
		if cerr := q.deleteCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryStmt: %w", cerr)
		}
	}
	if q.deleteStoragePlaceStmt != nil {
		if cerr := q.deleteStoragePlaceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteStoragePlaceStmt: %w", cerr)
		}
	}
	if q.getStorageConsumptionByIdStmt != nil {
		if cerr := q.getStorageConsumptionByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStorageConsumptionByIdStmt: %w", cerr)
		}
	}
	if q.getStorageItemByIdStmt != nil {
		if cerr := q.getStorageItemByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStorageItemByIdStmt: %w", cerr)
		}
	}
	if q.getStoragePlaceByIdStmt != nil {
		if cerr := q.getStoragePlaceByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStoragePlaceByIdStmt: %w", cerr)
		}
	}
	if q.listCategoriesStmt != nil {
		if cerr := q.listCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCategoriesStmt: %w", cerr)
		}
	}
	if q.listStorageConsumptionsStmt != nil {
		if cerr := q.listStorageConsumptionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStorageConsumptionsStmt: %w", cerr)
		}
	}
	if q.listStorageItemsStmt != nil {
		if cerr := q.listStorageItemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStorageItemsStmt: %w", cerr)
		}
	}
	if q.listStoragePlacesStmt != nil {
		if cerr := q.listStoragePlacesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStoragePlacesStmt: %w", cerr)
		}
	}
	if q.updateCategoryStmt != nil {
		if cerr := q.updateCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoryStmt: %w", cerr)
		}
	}
	if q.updateStorageItemStmt != nil {
		if cerr := q.updateStorageItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStorageItemStmt: %w", cerr)
		}
	}
	if q.updateStoragePlaceStmt != nil {
		if cerr := q.updateStoragePlaceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStoragePlaceStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	createCategoryStmt            *sql.Stmt
	createProductCategoryStmt     *sql.Stmt
	createStorageConsumptionStmt  *sql.Stmt
	createStorageItemStmt         *sql.Stmt
	createStoragePlaceStmt        *sql.Stmt
	deleteCategoryStmt            *sql.Stmt
	deleteStoragePlaceStmt        *sql.Stmt
	getStorageConsumptionByIdStmt *sql.Stmt
	getStorageItemByIdStmt        *sql.Stmt
	getStoragePlaceByIdStmt       *sql.Stmt
	listCategoriesStmt            *sql.Stmt
	listStorageConsumptionsStmt   *sql.Stmt
	listStorageItemsStmt          *sql.Stmt
	listStoragePlacesStmt         *sql.Stmt
	updateCategoryStmt            *sql.Stmt
	updateStorageItemStmt         *sql.Stmt
	updateStoragePlaceStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		createCategoryStmt:            q.createCategoryStmt,
		createProductCategoryStmt:     q.createProductCategoryStmt,
		createStorageConsumptionStmt:  q.createStorageConsumptionStmt,
		createStorageItemStmt:         q.createStorageItemStmt,
		createStoragePlaceStmt:        q.createStoragePlaceStmt,
		deleteCategoryStmt:            q.deleteCategoryStmt,
		deleteStoragePlaceStmt:        q.deleteStoragePlaceStmt,
		getStorageConsumptionByIdStmt: q.getStorageConsumptionByIdStmt,
		getStorageItemByIdStmt:        q.getStorageItemByIdStmt,
		getStoragePlaceByIdStmt:       q.getStoragePlaceByIdStmt,
		listCategoriesStmt:            q.listCategoriesStmt,
		listStorageConsumptionsStmt:   q.listStorageConsumptionsStmt,
		listStorageItemsStmt:          q.listStorageItemsStmt,
		listStoragePlacesStmt:         q.listStoragePlacesStmt,
		updateCategoryStmt:            q.updateCategoryStmt,
		updateStorageItemStmt:         q.updateStorageItemStmt,
		updateStoragePlaceStmt:        q.updateStoragePlaceStmt,
	}
}
