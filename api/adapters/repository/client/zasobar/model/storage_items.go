//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type StorageItems struct {
	StorageItemID  int32 `sql:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Title          *string
	StoragePlaceID *int32
	CategoryID     *int32
	BaselineAmount float64
	CurrentAmount  float64
	Quantity       StorageItemsQuantity
	Unit           string
	ExpirationDate *time.Time
	Ean            *string
}
