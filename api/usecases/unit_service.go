package usecases

import (
	goUnits "github.com/bcicen/go-units"
	"github.com/martinjirku/zasobar/domain"
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

func (u UnitService) ListAll() []domain.Unit {
	result := []domain.Unit{}
	for _, u := range goUnits.All() {
		if u.Quantity == "bytes" {
			continue
		}
		if u.Quantity == "bits" {
			continue
		}
		result = append(result, mapGoUnitsToDomain(u))
	}
	return result
}

func (u UnitService) ListByQuantity(quantity domain.Quantity) ([]domain.Unit, error) {
	var units = []domain.Unit{}
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

func mapGoUnitsToDomain(u goUnits.Unit) domain.Unit {
	return domain.Unit{
		Name:       u.Name,
		Symbol:     u.Symbol,
		Quantity:   u.Symbol,
		PluralName: u.PluralName(),
		Names:      u.Names(),
		System:     u.System(),
	}
}
