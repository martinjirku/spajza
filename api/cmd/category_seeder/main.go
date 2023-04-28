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
	opt := options{}
	flag.StringVar(&opt.filePath, "file", "./data/taxonomy-with-ids.sk-SK.txt", "path for file")
	flag.StringVar(&opt.db.Name, "name", "zasobar", "db name")
	flag.StringVar(&opt.db.User, "user", "user", "db name")
	flag.StringVar(&opt.db.Host, "host", "localhost", "db host")
	flag.StringVar(&opt.db.Port, "port", "3306", "db port")
	flag.StringVar(&opt.db.Password, "pwd", "", "db pwd")
	flag.StringVar(&opt.db.Type, "type", "mysql", "db type")

	flag.Parse()

	readFile, err := os.Open("./data/taxonomy-with-ids.sk-SK.txt")

	defer func() { readFile.Close() }()

	if err != nil {
		panic("could not open the file")
	}

	database := db.NewDB(opt.db)

	tx, err := database.BeginTx(context.Background(), nil)
	defer tx.Rollback()
	if err != nil {
		log.Fatalf("could not create transaction %s", err)
	}
	cats := productCategory.GetProductCategories(readFile)
	categories := []entity.ProductCategory{}

	categoryIndex := map[entity.ProducCategoryPath]*int64{}
	for _, c := range cats {
		category := entity.ProductCategory{
			CategoryId: c.Id,
			Name:       c.Name,
			Path:       entity.NewProductCategoryPath(c.Path),
		}
		categories = append(categories, category)
		if category.Path == "" || len(c.Path) == 0 {
			continue
		}
		parentPath := c.Path[0:len(c.Path)]

		categoryIndex[entity.NewProductCategoryPath(parentPath)] = &c.Id

	}

	for _, c := range categories {
		if _, ok := categoryIndex[c.Path]; !ok && c.Path != "" {
			log.Default().Printf("Could not find parent category ID for %v", c.CategoryId)
		}
		c.ParentId = categoryIndex[c.Path]
		_, err := tx.Exec("INSERT INTO product_categories(category_id,name,path,parent_id) VALUES (?,?,?,?)",
			c.CategoryId, c.Name, c.Path, c.ParentId)
		if err != nil {
			log.Default().Printf("Could not insert product_category %v: %q", c.CategoryId, c.Name)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("could not commit transaction %s", err)
	}
}

type options struct {
	filePath string
	db       config.Db
}
