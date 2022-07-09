package units

import (
	goUnits "github.com/bcicen/go-units"
)

type UnitService struct {
}

var isInitialized = false

func NewUnitService() UnitService {
	if isInitialized {
		initUnits()
	}
	return UnitService{}
}

func initUnits() {
	goUnits.NewUnit("cup", "cup", goUnits.Volume)
	isInitialized = true
}

func (u UnitService) ListAll() []goUnits.Unit {
	result := []goUnits.Unit{}
	for _, u := range goUnits.All() {
		if u.Quantity == "bytes" {
			continue
		}
		if u.Quantity == "bits" {
			continue
		}
		result = append(result, u)
	}
	return result
}

func (u UnitService) ListByQuantity(quantity Quantity) ([]goUnits.Unit, error) {
	var units = []goUnits.Unit{}
	for _, unit := range u.ListAll() {
		q, err := quantity.Value()
		if err != nil {
			return units, err
		}
		if q == unit.Quantity {
			units = append(units, unit)
		}
	}
	return units, nil
}
