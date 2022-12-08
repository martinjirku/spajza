package entity

import (
	"strconv"
	"strings"
)

type CategoryPath string

type Category struct {
	ID          int64
	Title       string
	Path        CategoryPath
	DefaultUnit string
}

func (c *Category) Validate() bool {
	if c.Title == "" {
		return false
	}
	if c.DefaultUnit == "" {
		return false
	}
	return true
}

func (c *Category) hasSameTitle(title string) bool {
	return strings.EqualFold(c.Title, title)
}

func (c *Category) validateAgainstCategory(category Category) error {
	if c.hasSameTitle(category.Title) {
		return ErrEntityConflict
	}
	return nil
}

type Categories []Category

func (c *Categories) SetCategoryTitle(id int64, title string) (Category, error) {
	result := Category{}
	categories := *c
	var category *Category
	for i := 0; i < len(categories); i++ {
		if categories[i].ID == id {
			category = &categories[i]
			break
		}
		if strings.EqualFold(categories[i].Title, title) {
			return result, ErrEntityConflict
		}
	}
	if category == nil {
		return result, ErrEntityNotFound
	}
	category.Title = title
	result = *category
	return result, nil
}

func (c *Categories) AddCategory(category Category) error {
	if !category.Validate() {
		return ErrInvalidEntity
	}
	err := c.validateCategory(category)
	if err != nil {
		return err
	}
	*c = append(*c, category)
	return nil
}

func (c *Categories) validateCategory(category Category) error {
	categories := *c
	categoryMap := map[int64]*Category{}
	for i := 0; i < len(categories); i++ {
		if err := categories[i].validateAgainstCategory(category); err != nil {
			return err
		}
		categoryMap[categories[i].ID] = &categories[i]
	}
	// validate Path
	path := string(category.Path)
	sections := strings.Split(path, ".")
	for idx, section := range sections {
		if idx == 0 && section == "" {
			continue
		}
		id, err := strconv.ParseInt(section, 10, 64)
		if err != nil {
			return ErrInvalidEntity
		}
		if _, ok := categoryMap[id]; !ok {
			return ErrInvalidEntity
		}
	}
	return nil
}
