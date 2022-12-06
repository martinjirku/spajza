package entity

import (
	"math"
	"time"

	u "github.com/bcicen/go-units"
)

type StoragePlace struct {
	StoragePlaceId uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Code           string
}

type NewStorageItem struct {
	CategoryId     uint      `json:"categoryId"`
	StoragePlaceId uint      `json:"storagePlaceId"`
	Title          string    `json:"title"`
	Amount         float64   `json:"amount"`
	Unit           string    `json:"unit"`
	ExpirationDate time.Time `json:"expirationDate"`
}

// StorageItem is struct to track item in storage and its consumption
type StorageItem struct {
	StorageItemId   uint                     `json:"storageItemId"`
	Title           string                   `json:"title"`
	BaselineAmount  float64                  `json:"baselineAmount"`
	CurrentAmount   float64                  `json:"currentAmount"`
	CategoryId      uint                     `json:"categoryId"`
	StoragePlaceId  uint                     `json:"storagePlaceId"`
	StorageLocation string                   `json:"storageLocation"`
	Quantity        Quantity                 `json:"quantity"`
	Unit            string                   `json:"unit"`
	ExpirationDate  time.Time                `json:"expirationDate"`
	Consumptions    []StorageItemConsumption `json:"consumptions,omitempty"`
}

type StorageItemConsumption struct {
	StorageItemConsumptionId uint
	NormalizedAmount         float64
	Unit                     string
	StorageItemId            uint
}

type StorageItemLoader interface {
	GetStorageItemById(storageItemId uint) (StorageItem, error)
	GetStorageConsumptionById(storageItemId uint) ([]StorageItemConsumption, error)
}

// "LoadStorageItem" loads StorageItem from database.
//
// One needs to provide storageItemId, context and loader.
func LoadStorageItem(id uint, loader StorageItemLoader) (StorageItem, error) {
	storageItem, err := loader.GetStorageItemById(id)
	if err != nil {
		return storageItem, err
	}
	consumptions, err := loader.GetStorageConsumptionById(id)
	if err != nil {
		return storageItem, err
	}
	storageItem.Consumptions = consumptions
	return storageItem, nil
}

// When Consumpt is called, specified amount will be removed
// from the current amount.
func (s *StorageItem) Consumpt(amount float64, un string) error {
	normalizedAmount := amount
	if s.Unit != un {
		itemUnit, err := u.Find(s.Unit)
		if err != nil {
			return err
		}
		consumptedUnit, err := u.Find(un)
		if err != nil {
			return err
		}
		normalizedAmountValue, err := u.ConvertFloat(amount, consumptedUnit, itemUnit)
		if err != nil {
			return err
		}
		normalizedAmount = normalizedAmountValue.Float()
	}
	consumption := StorageItemConsumption{NormalizedAmount: normalizedAmount, StorageItemId: s.StorageItemId, Unit: un}
	s.Consumptions = append(s.Consumptions, consumption)

	currentAmount := s.CurrentAmount - normalizedAmount
	s.CurrentAmount = math.Max(currentAmount, 0)
	return nil
}
