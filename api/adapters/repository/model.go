package repository

import "time"

type StorageItem struct {
	StorageItemId  uint
	Title          string
	BaselineAmount float64
	CurrentAmount  float64
	CategoryId     uint
	StoragePlaceId uint
	QuantityType   string
	Unit           string
	Ean            string
	ExpirationDate time.Time
}
