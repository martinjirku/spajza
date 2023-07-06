package entity

import (
	"time"
)

type StoragePlace struct {
	StoragePlaceId int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string
	Code           string
}

type StorageItem struct {
	StorageItemID    int32
	Title            string
	CategoryID       int32
	StoragePlaceID   int32
	Ean              string
	ExpirationDate   time.Time
	baselineQuantity Quantity
	consumptions     []StorageItemConsumption
}

type StorageItemList struct {
	Items []StorageItem `json:"items"`
	Count int64         `json:"count"`
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

func (s *StorageItem) AddConsumption(consumption StorageItemConsumption) error {
	if consumption.Quantity.Unit.GetQuantityType() != s.baselineQuantity.Unit.GetQuantityType() {
		return ErrInvalidParameter
	}
	err := consumption.Quantity.Verify()
	if err != nil {
		return err
	}
	s.consumptions = append(s.consumptions, consumption)
	return nil
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
	result, err := s.CurrentQuantity().Subtract(consumption.Quantity)
	if err != nil {
		return err
	}
	if result.Value < 0. {
		return ErrStorageItemNotEnoughQuantity
	}
	err = consumption.Quantity.Verify()
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
	StorageItemConsumptionID int32
	Quantity                 Quantity
}
