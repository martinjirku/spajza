package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type BulkProductCategoryCreator interface {
	CreateCategories(categories []entity.ProductCategory) error
}

type ProductCategoryUsecase struct {
	bulkCreator BulkProductCategoryCreator
}

func NewProductCategoryUsecase(bulkCreator BulkProductCategoryCreator) ProductCategoryUsecase {
	return ProductCategoryUsecase{bulkCreator: bulkCreator}
}

// func (pc *ProductCategoryUsecase) BulkCreate(rd io.Reader) error {
// 	// productCategories, err := productCategory.Read(rd)
// 	if err != nil {
// 		return err
// 	}
// 	categoryMap := map[string]int64{}
// 	categories := make([]entity.ProductCategory, len(productCategories))
// 	for i, c := range productCategories {
// 		name := ""
// 		for idx := len(c.Name) - 1; idx > 0; idx-- {
// 			if c.Name[idx] == '>' {
// 				name = c.Name[idx+1:]
// 				break
// 			}

// 		}
// 		categories[i] = entity.ProductCategory{
// 			CategoryId: c.Id,
// 			Path:       c.Name,
// 			Name:       strings.TrimFunc(name, unicode.IsSpace),
// 		}
// 		categoryMap[categories[i].Name] = c.Id
// 	}
// 	// for _, c := range categories {

// 	// }
// 	return pc.bulkCreator.CreateCategories(categories)
// }
