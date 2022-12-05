package usecase

import (
	goUnits "github.com/bcicen/go-units"
	"github.com/martinjirku/zasobar/entity"
)

type UnitUsecase struct{}

var isInitialized = false
var supportedUnits = []string{
	"gram", "milligram", "kilogram", "decagram", "pound", "ounce", // mass
	"meter", "centimeter", "decimeter", "foot", "inch", "kilometer", "mile", "yard", // length
	"gallon", "hectoliter", "liter", "milliliter", "pint", // volume
	"celsius", "fahrenheit", "kelvin", // temperature
	"century", "day", "decade", "hour", "millisecond", "minute", "month", "year", // time
	"count", // count
}

func NewUnitService() UnitUsecase {
	if isInitialized {
		initUnits()
	}
	return UnitUsecase{}
}

func initUnits() {
	goUnits.NewUnit("cup", "cup", goUnits.Volume)
	isInitialized = true
}

func (u UnitUsecase) ListAll() []entity.Unit {
	result := []entity.Unit{}
	for _, unit := range supportedUnits {
		u, err := goUnits.Find(unit)
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

func (u UnitUsecase) ListByQuantity(quantity entity.Quantity) ([]entity.Unit, error) {
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
		Name:       u.Name,
		Symbol:     u.Symbol,
		Quantity:   entity.Quantity(u.Quantity),
		PluralName: u.PluralName(),
		Names:      u.Names(),
		System:     u.System(),
	}
}
