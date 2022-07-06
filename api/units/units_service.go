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
	return goUnits.All()
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
