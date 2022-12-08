package usecase

import (
	goUnits "github.com/bcicen/go-units"
	"github.com/martinjirku/zasobar/entity"
)

type UnitUsecase struct{}

var isInitialized = false

func NewUnitUsecase() UnitUsecase {
	if isInitialized {
		initUnits()
	}
	return UnitUsecase{}
}

func initUnits() {
	goUnits.NewUnit("cup", "cup", goUnits.Volume)
	isInitialized = true
}

func (u *UnitUsecase) ListAll() []entity.Unit {
	result := []entity.Unit{}
	for _, unit := range entity.SupportedUnits {
		u, err := goUnits.Find(string(unit))
		if err != nil {
			if unit == "count" {
				result = append(result, entity.Unit{
					Name:     "count",
					Quantity: "count",
					Symbol:   "ks",
				})
			}
			continue
		}
		result = append(result, mapGoUnitsToDomain(u))
	}
	return result
}

func (u *UnitUsecase) ListByQuantity(quantity entity.Quantity) ([]entity.Unit, error) {
	var units = []entity.Unit{}
	for _, unit := range u.ListAll() {
		q, err := quantity.Value()
		if err != nil {
			return units, err
		}
		if entity.Quantity(q) == unit.Quantity {
			units = append(units, unit)
		}
	}
	return units, nil
}

func mapGoUnitsToDomain(u goUnits.Unit) entity.Unit {
	return entity.Unit{
		Name:       entity.UnitName(u.Name),
		Symbol:     u.Symbol,
		Quantity:   entity.Quantity(u.Quantity),
		PluralName: u.PluralName(),
		Names:      u.Names(),
		System:     u.System(),
	}
}
