package web

import (
	"net/http"

	goUnits "github.com/bcicen/go-units"
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/units"
)

type (
	unit struct {
		Name       string   `json:"name"`
		Names      []string `json:"names"`
		PluralName string   `json:"pluralName"`
		Symbol     string   `json:"symbol"`
		System     string   `json:"system"`
		Quantity   string   `json:"quantity"`
	}
)

func mapUnitToDto(u goUnits.Unit) unit {

	return unit{
		Name:       u.Name,
		Quantity:   u.Quantity,
		Symbol:     u.Symbol,
		System:     u.System(),
		Names:      u.Names(),
		PluralName: u.PluralName(),
	}
}

func mapGoUnitsToUnits(u []goUnits.Unit) []unit {
	var units = []unit{}
	for _, unit := range u {
		units = append(units, mapUnitToDto(unit))
	}
	return units
}

func listHandler(c echo.Context) error {
	unitService := units.NewUnitService()
	return c.JSON(http.StatusOK, mapGoUnitsToUnits(unitService.ListAll()))
}

func listUnitsByQuantityHandler(c echo.Context) error {
	unitService := units.NewUnitService()
	var quantity units.Quantity
	err := quantity.Scan(c.Param("quantity"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	units, err := unitService.ListByQuantity(quantity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mapGoUnitsToUnits(units))
}
