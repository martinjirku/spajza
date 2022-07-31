package domain

import "time"

type StoragePlace struct {
	StoragePlaceId uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Code           string
}

type StorageItem struct {
	StorageItemId   uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Title           string
	BaselineAmount  float32
	CurrentAmount   float32
	CategoryId      uint
	StoragePlaceId  uint
	StorageLocation string
	Quantity        Quantity
	ExpirationDate  time.Time
}

type NewStorageItemRequest struct {
	CategoryId     uint      `json:"categoryId"`
	StoragePlaceId uint      `json:"storagePlaceId"`
	Title          string    `json:"title"`
	Amount         float32   `json:"amount"`
	Unit           string    `json:"unit"`
	ExpirationDate time.Time `json:"expirationDate"`
}

type StorageItemResponse struct {
	StorageItemId   uint      `json:"storageItemId"`
	Title           string    `json:"title"`
	BaselineAmount  float32   `json:"baselineAmount"`
	CurrentAmount   float32   `json:"currentAmount"`
	CategoryId      uint      `json:"categoryId"`
	StoragePlaceId  uint      `json:"storagePlaceId"`
	StorageLocation string    `json:"storageLocation"`
	Quantity        Quantity  `json:"quantity"`
	Unit            string    `json:"unit"`
	ExpirationDate  time.Time `json:"expirationDate"`
}
