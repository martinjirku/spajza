package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/infra/db"
	"github.com/martinjirku/zasobar/pkg/google/productCategory"
)

func main() {
	var filePath string
	var dbConf config.Db

	flag.StringVar(&filePath, "file", "./cmd/category_seeder/data/taxonomy-with-ids.sk-SK.txt", "path for file")
	flag.StringVar(&dbConf.Name, "name", "zasobar", "db name")
	flag.StringVar(&dbConf.User, "user", "user", "db name")
	flag.StringVar(&dbConf.Host, "host", "localhost", "db host")
	flag.StringVar(&dbConf.Port, "port", "3306", "db port")
	flag.StringVar(&dbConf.Password, "pwd", "", "db pwd")
	flag.StringVar(&dbConf.Type, "type", "mysql", "db type")

	flag.Parse()

	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not Open faile: %v", err)
	}
	defer func() { readFile.Close() }()

	database := db.NewDB(dbConf)

	tx, err := database.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		log.Fatalf("could not create transaction %s", err)
	}
	cats := productCategory.GetProductCategories(readFile)
	categories := []entity.ProductCategory{}

	categoryIndex := map[entity.ProducCategoryPath]*int32{}
	for _, c := range cats {
		category := entity.ProductCategory{
			CategoryId: int32(c.Id),
			Name:       c.Name,
			Path:       entity.NewProductCategoryPath(c.Path),
		}
		category.ParentId = categoryIndex[category.Path]

		categories = append(categories, category)
		parentPath := c.Path[0:len(c.Path)]
		path := entity.NewProductCategoryPath(append(parentPath, category.Name))
		categoryIndex[path] = &category.CategoryId
	}

	for _, c := range categories {
		if _, ok := categoryIndex[c.Path]; !ok && c.Path != "" {
			log.Default().Printf("Could not find parent category ID for %v", c.CategoryId)
		}
		c.ParentId = categoryIndex[c.Path]
		_, err := tx.Exec("INSERT INTO product_categories(category_id,name,path,parent_id) VALUES (?,?,?,?)",
			c.CategoryId, c.Name, c.Path, c.ParentId)
		if err != nil {
			log.Default().Printf("Could not insert product_category %v: %q: %q", c.CategoryId, c.Name, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("could not commit transaction %s", err)
	}
}
