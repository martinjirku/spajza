package repository

import "time"

type StorageItem struct {
	StorageItemId  int32
	Title          string
	BaselineAmount float64
	CurrentAmount  float64
	CategoryId     int32
	StoragePlaceId int32
	QuantityType   string
	Unit           string
	Ean            string
	ExpirationDate time.Time
}
