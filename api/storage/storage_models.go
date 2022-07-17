package storage

import "time"

type StoragePlace struct {
	StoragePlaceId uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Code           string
}
