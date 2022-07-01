package services

import (
	goUnits "github.com/bcicen/go-units"
)

type UnitService struct {
}

var isInitialized = false

func NewUnitService() UnitService {
	if isInitialized {
		InitUnits()
	}
	return UnitService{}
}

func InitUnits() {
	goUnits.NewUnit("cup", "cup", goUnits.Volume)
	isInitialized = true
}

func (u UnitService) ListAll() []goUnits.Unit {
	return goUnits.All()
}

func (u UnitService) ListByQuantity(quantity string) []goUnits.Unit {
	var units = []goUnits.Unit{}
	for _, unit := range u.ListAll() {
		if quantity == unit.Quantity {
			units = append(units, unit)
		}
	}
	return units
}
