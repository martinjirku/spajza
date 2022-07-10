package categories

import (
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return CategoryService{db: db}
}

func (cs *CategoryService) ListAll() ([]Category, error) {
	categories := new([]Category)
	result := cs.db.Find(categories)
	if result.Error != nil {
		return *categories, result.Error
	}
	return *categories, nil
}

func (cs *CategoryService) CreateItem(c Category) (Category, error) {
	result := cs.db.Create(&c)
	return c, result.Error
}

func (cs *CategoryService) UpdateItem(c Category) (Category, error) {
	result := cs.db.Model(c).Updates(Category{Title: c.Title, Path: c.Path, DefaultUnit: c.DefaultUnit})
	return c, result.Error

}
func (cs *CategoryService) DeleteItem(c Category) error {
	result := cs.db.Delete(&c)
	return result.Error
}
