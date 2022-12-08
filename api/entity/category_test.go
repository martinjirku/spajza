package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func generateCategories() entity.Categories {
	return entity.Categories{
		{1, "Title1", "", "gram"},
		{2, "Title2", "", "gram"},
	}
}

func Test_Categories_SetCategoryTitle(t *testing.T) {
	t.Run("Set correct non duplicated title", func(t *testing.T) {
		categories := generateCategories()
		c, err := categories.SetCategoryTitle(2, "Title3")
		if err != nil {
			t.Errorf("Expected nil, given %s", err)
		}
		if c.Title != "Title3" && categories[1].Title == "Title3" {
			t.Errorf("Expected Title3, given %s", c.Title)
		}
	})
	t.Run("Set duplicate title", func(t *testing.T) {
		categories := entity.Categories{
			{1, "Title1", "", "gram"},
			{2, "Title2", "", "gram"},
		}
		_, err := categories.SetCategoryTitle(2, "Title1")
		if err != entity.ErrEntityConflict {
			t.Errorf("Expected %s, given %s", entity.ErrEntityConflict, err)
		}
		if categories[1].Title != "Title2" {
			t.Errorf("Title can not be changed, when there is duplicity")
		}
	})
	t.Run("Set duplicate caseinsensitive title", func(t *testing.T) {
		categories := generateCategories()
		_, err := categories.SetCategoryTitle(2, "tItle1")
		if err != entity.ErrEntityConflict {
			t.Errorf("Expected %s, given %s", entity.ErrEntityConflict, err)
		}
		if categories[1].Title != "Title2" {
			t.Errorf("Title can not be changed, when there is duplicity")
		}
	})
	t.Run("Set non existing category Title", func(t *testing.T) {
		categories := generateCategories()
		_, err := categories.SetCategoryTitle(3, "tutle1")
		if err != entity.ErrEntityNotFound {
			t.Errorf("Expected %s, given %s", entity.ErrEntityNotFound, err)
		}
	})
}

func Test_AddCategory(t *testing.T) {
	t.Run("Add correct category", func(t *testing.T) {
		categories := generateCategories()
		category := entity.Category{0, "Special Category", entity.CategoryPath("1"), "gram"}
		err := categories.AddCategory(category)
		if err != nil {
			t.Errorf("Expected no error, given %s", err)
		}
		if len(categories) != 3 {
			t.Errorf("Expected length %d, given %d", 3, len(categories))
		}
		if categories[2].Title != "Special Category" {
			t.Errorf("Expected %s, given %s", "Special Category", categories[2].Title)
		}
	})
	t.Run("Add category with duplicate title", func(t *testing.T) {
		categories := generateCategories()
		category := entity.Category{0, "Title1", entity.CategoryPath("1"), "gram"}
		err := categories.AddCategory(category)
		if err == nil {
			t.Error("Expected error, but no error received")
		} else if err != entity.ErrEntityConflict {
			t.Errorf("Expected %s, given %s", entity.ErrEntityConflict, err)
		}
	})
	t.Run("Add category with wrong path", func(t *testing.T) {
		categories := generateCategories()
		category := entity.Category{0, "Special Category", entity.CategoryPath("3"), "gram"}
		err := categories.AddCategory(category)
		if err == nil {
			t.Error("Expected error, but no error received")
		} else if err != entity.ErrInvalidEntity {
			t.Errorf("Expected %s, given %s", entity.ErrInvalidEntity, err)
		}
	})
	t.Run("Add category with correct path", func(t *testing.T) {
		categories := generateCategories()
		path := ".2"
		category := entity.Category{0, "Special Category", entity.CategoryPath(path), "gram"}
		err := categories.AddCategory(category)
		if err != nil {
			t.Error("Expected no error, but error received")
		}

		if category.Path != entity.CategoryPath(path) && categories[2].Path != entity.CategoryPath(path) {
			t.Errorf("Expected %s, but received %s", path, category.Path)
		}
	})
}

func Test_CategoryValidate(t *testing.T) {
	t.Run("Title cannot be empty", func(t *testing.T) {
		category := entity.Category{
			ID:          1,
			Title:       "",
			Path:        "Path1",
			DefaultUnit: "gram",
		}
		if category.Validate() {
			t.Error("Title is invalid, but it returned as valid")
		}
	})
	t.Run("Default Unit cannot be empty", func(t *testing.T) {
		category := entity.Category{
			ID:          1,
			Title:       "",
			Path:        "Path1",
			DefaultUnit: "gram",
		}
		if category.Validate() {
			t.Error("Default unit is invalid, but it returned as valid")
		}
	})
}
