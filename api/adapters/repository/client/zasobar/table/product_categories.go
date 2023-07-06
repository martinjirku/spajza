//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var ProductCategories = newProductCategoriesTable("zasobar", "product_categories", "")

type productCategoriesTable struct {
	mysql.Table

	// Columns
	CategoryID mysql.ColumnInteger
	Name       mysql.ColumnString
	Path       mysql.ColumnString
	ParentID   mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ProductCategoriesTable struct {
	productCategoriesTable

	NEW productCategoriesTable
}

// AS creates new ProductCategoriesTable with assigned alias
func (a ProductCategoriesTable) AS(alias string) *ProductCategoriesTable {
	return newProductCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ProductCategoriesTable with assigned schema name
func (a ProductCategoriesTable) FromSchema(schemaName string) *ProductCategoriesTable {
	return newProductCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ProductCategoriesTable with assigned table prefix
func (a ProductCategoriesTable) WithPrefix(prefix string) *ProductCategoriesTable {
	return newProductCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ProductCategoriesTable with assigned table suffix
func (a ProductCategoriesTable) WithSuffix(suffix string) *ProductCategoriesTable {
	return newProductCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newProductCategoriesTable(schemaName, tableName, alias string) *ProductCategoriesTable {
	return &ProductCategoriesTable{
		productCategoriesTable: newProductCategoriesTableImpl(schemaName, tableName, alias),
		NEW:                    newProductCategoriesTableImpl("", "new", ""),
	}
}

func newProductCategoriesTableImpl(schemaName, tableName, alias string) productCategoriesTable {
	var (
		CategoryIDColumn = mysql.IntegerColumn("category_id")
		NameColumn       = mysql.StringColumn("name")
		PathColumn       = mysql.StringColumn("path")
		ParentIDColumn   = mysql.IntegerColumn("parent_id")
		allColumns       = mysql.ColumnList{CategoryIDColumn, NameColumn, PathColumn, ParentIDColumn}
		mutableColumns   = mysql.ColumnList{NameColumn, PathColumn, ParentIDColumn}
	)

	return productCategoriesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CategoryID: CategoryIDColumn,
		Name:       NameColumn,
		Path:       PathColumn,
		ParentID:   ParentIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
