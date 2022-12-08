package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/martinjirku/zasobar/adapters/repository"
)

func Test_RepositoryCategoryListAll(t *testing.T) {
	query := "SELECT id, title, default_unit, path FROM categories WHERE deleted_at IS null"
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.
			ExpectQuery(query).WillReturnRows(sqlmock.
			NewRows([]string{"id", "title", "default_unit", "path"}).
			FromCSVString("1,\"title1\",\"gram\",\n2,\"title2\",\"gram\",\"1\""))

		categoryRepo := repository.NewCategoryRepository(context.Background(), db)
		allItems, err := categoryRepo.List()
		if err != nil {
			t.Errorf("ListAll should not return an error %s", err)
		}
		if len(allItems) != 2 {
			t.Errorf("ListAll should return two items, but it return %d", len(allItems))
		}
		if allItems[0].DefaultUnit != "gram" {
			t.Errorf("Expected %s, but received %s", "gram", allItems[0].DefaultUnit)
		}
		if allItems[0].Path != "" {
			t.Errorf("Expected %s, but received %s", "", allItems[0].Path)
		}
		if allItems[0].Title != "title1" {
			t.Errorf("Expected %s, but received %s", "title1", allItems[0].Title)
		}
	})
	t.Run("NoResults", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "default_unit", "path"}))

		categoryRepo := repository.NewCategoryRepository(context.Background(), db)
		allItems, err := categoryRepo.List()
		if err != nil {
			t.Errorf("ListAll should not return an error %s", err)
		}
		if len(allItems) != 0 {
			t.Errorf("ListAll should return two items, but it return %d", len(allItems))
		}
	})
}
