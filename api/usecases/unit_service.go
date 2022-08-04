package usecases

import (
	goUnits "github.com/bcicen/go-units"
	"github.com/martinjirku/zasobar/domain"
)

type UnitService struct {
}

var isInitialized = false
var supportedUnits = []string{
	"gram", "milligram", "kilogram", "decagram", "pound", "ounce", // mass
	"meter", "centimeter", "decimeter", "foot", "inch", "kilometer", "mile", "yard", // length
	"gallon", "hectoliter", "liter", "milliliter", "pint", // volume
	"celsius", "fahrenheit", "kelvin", // temperature
	"century", "day", "decade", "hour", "millisecond", "minute", "month", "year", // time
	"count", // count
}

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
	for _, unit := range supportedUnits {
		u, err := goUnits.Find(unit)
		if err != nil {
			if unit == "count" {
				result = append(result, domain.Unit{
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

func (u UnitService) ListByQuantity(quantity domain.Quantity) ([]domain.Unit, error) {
	var units = []domain.Unit{}
	for _, unit := range u.ListAll() {
		q, err := quantity.Value()
		if err != nil {
			return units, err
		}
		if domain.Quantity(q) == unit.Quantity {
			units = append(units, unit)
		}
	}
	return units, nil
}

func mapGoUnitsToDomain(u goUnits.Unit) domain.Unit {
	return domain.Unit{
		Name:       u.Name,
		Symbol:     u.Symbol,
		Quantity:   domain.Quantity(u.Quantity),
		PluralName: u.PluralName(),
		Names:      u.Names(),
		System:     u.System(),
	}
}
