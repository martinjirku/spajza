package web

import (
	"github.com/martinjirku/zasobar/entity"
)

type (
	unitDto struct {
		Name       string          `json:"name"`
		Names      []string        `json:"names"`
		PluralName string          `json:"pluralName"`
		Symbol     string          `json:"symbol"`
		System     string          `json:"system"`
		Quantity   entity.Quantity `json:"quantity"`
	}
)

func UnitDto(u entity.Unit) unitDto {
	return unitDto{
		Name:       u.Name,
		Quantity:   u.Quantity,
		Symbol:     u.Symbol,
		System:     u.System,
		Names:      u.Names,
		PluralName: u.PluralName,
	}
}

func mapGoUnitsToUnitDto(u []entity.Unit) []unitDto {
	var units = make([]unitDto, len(u))
	for i, unit := range u {
		units[i] = UnitDto(unit)
	}
	return units
}
