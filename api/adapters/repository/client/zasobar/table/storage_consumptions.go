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

var StorageConsumptions = newStorageConsumptionsTable("zasobar", "storage_consumptions", "")

type storageConsumptionsTable struct {
	mysql.Table

	// Columns
	StorageItemConsumptionID mysql.ColumnInteger
	CreatedAt                mysql.ColumnTimestamp
	UpdatedAt                mysql.ColumnTimestamp
	DeletedAt                mysql.ColumnTimestamp
	NormalizedAmount         mysql.ColumnFloat
	Unit                     mysql.ColumnString
	StorageItemID            mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type StorageConsumptionsTable struct {
	storageConsumptionsTable

	NEW storageConsumptionsTable
}

// AS creates new StorageConsumptionsTable with assigned alias
func (a StorageConsumptionsTable) AS(alias string) *StorageConsumptionsTable {
	return newStorageConsumptionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StorageConsumptionsTable with assigned schema name
func (a StorageConsumptionsTable) FromSchema(schemaName string) *StorageConsumptionsTable {
	return newStorageConsumptionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StorageConsumptionsTable with assigned table prefix
func (a StorageConsumptionsTable) WithPrefix(prefix string) *StorageConsumptionsTable {
	return newStorageConsumptionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StorageConsumptionsTable with assigned table suffix
func (a StorageConsumptionsTable) WithSuffix(suffix string) *StorageConsumptionsTable {
	return newStorageConsumptionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStorageConsumptionsTable(schemaName, tableName, alias string) *StorageConsumptionsTable {
	return &StorageConsumptionsTable{
		storageConsumptionsTable: newStorageConsumptionsTableImpl(schemaName, tableName, alias),
		NEW:                      newStorageConsumptionsTableImpl("", "new", ""),
	}
}

func newStorageConsumptionsTableImpl(schemaName, tableName, alias string) storageConsumptionsTable {
	var (
		StorageItemConsumptionIDColumn = mysql.IntegerColumn("storage_item_consumption_id")
		CreatedAtColumn                = mysql.TimestampColumn("created_at")
		UpdatedAtColumn                = mysql.TimestampColumn("updated_at")
		DeletedAtColumn                = mysql.TimestampColumn("deleted_at")
		NormalizedAmountColumn         = mysql.FloatColumn("normalized_amount")
		UnitColumn                     = mysql.StringColumn("unit")
		StorageItemIDColumn            = mysql.IntegerColumn("storage_item_id")
		allColumns                     = mysql.ColumnList{StorageItemConsumptionIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NormalizedAmountColumn, UnitColumn, StorageItemIDColumn}
		mutableColumns                 = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NormalizedAmountColumn, UnitColumn, StorageItemIDColumn}
	)

	return storageConsumptionsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StorageItemConsumptionID: StorageItemConsumptionIDColumn,
		CreatedAt:                CreatedAtColumn,
		UpdatedAt:                UpdatedAtColumn,
		DeletedAt:                DeletedAtColumn,
		NormalizedAmount:         NormalizedAmountColumn,
		Unit:                     UnitColumn,
		StorageItemID:            StorageItemIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
