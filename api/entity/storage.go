package entity

import (
	"time"
)

type StoragePlace struct {
	StoragePlaceId uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Code           string
}

// StorageItem is struct to track item in storage and its consumption
type StorageItem struct {
	StorageItemId    uint
	Title            string
	CategoryId       uint
	StoragePlaceId   uint
	ExpirationDate   time.Time
	baselineQuantity Quantity
	consumptions     []StorageItemConsumption
}

func (s *StorageItem) Init() *StorageItem {
	if s.consumptions == nil {
		s.consumptions = []StorageItemConsumption{}
	}
	return s
}

func (s *StorageItem) Validate() error {
	if validateConsumption(s.consumptions, s.baselineQuantity.Unit.GetQuantityType()) {
		return ErrInvalidEntity
	}
	return nil
}

func (s *StorageItem) BaselineQuantity() Quantity {
	return s.baselineQuantity
}

func (s *StorageItem) SetBaselineQuantity(q Quantity) error {
	if !validateConsumption(s.consumptions, q.Unit.GetQuantityType()) {
		return ErrInvalidParameter
	}
	s.baselineQuantity = q
	return nil
}

func (s *StorageItem) Consumptions() []StorageItemConsumption {
	return s.consumptions
}

func (s *StorageItem) SetConsumptions(consumptions []StorageItemConsumption) error {
	s.consumptions = consumptions
	return nil
}

func (s *StorageItem) Consumpt(v float64, u UnitName) error {
	if u.GetQuantityType() != s.baselineQuantity.Unit.GetQuantityType() {
		return ErrInvalidParameter
	}
	consumption := StorageItemConsumption{Quantity: Quantity{v, u}}
	err := consumption.Quantity.Verify()
	if err != nil {
		return err
	}
	s.consumptions = append(s.consumptions, consumption)
	return nil
}

func (s *StorageItem) CurrentQuantity() Quantity {
	result := s.baselineQuantity
	for _, c := range s.consumptions {
		result, _ = result.Subtract(c.Quantity)
	}
	return result
}

func validateConsumption(consumptions []StorageItemConsumption, quantityType QuantityType) bool {
	for _, c := range consumptions {
		if c.Quantity.Unit.GetQuantityType() != quantityType {
			return false
		}
	}
	return true
}

type StorageItemConsumption struct {
	StorageItemConsumptionId uint
	Quantity                 Quantity
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
	if consumptions == nil {
		consumptions = []StorageItemConsumption{}
	}
	storageItem.consumptions = consumptions
	return storageItem, nil
}
