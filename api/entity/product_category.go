package entity

import "strings"

type ProducCategoryPath string

func NewProductCategoryPath(paths []string) ProducCategoryPath {
	return ProducCategoryPath(strings.Join(paths, " > "))
}

type ProductCategory struct {
	CategoryId int64              `json:"categoryId"`
	Path       ProducCategoryPath `json:"path"`
	Name       string             `json:"name"`
	ParentId   *int64             `json:"parentId"`
}
