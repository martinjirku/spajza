package units

import (
	"net/http"

	goUnits "github.com/bcicen/go-units"
	"github.com/labstack/echo/v4"
)

type (
	unitController struct {
		unitService UnitService
	}

	unit struct {
		Name     string
		Symbol   string
		Type     string
		Quantity string
	}
)

func newUnitController() unitController {
	return unitController{unitService: NewUnitService()}
}

func mapUnitToDto(u goUnits.Unit) unit {
	return unit{
		Name:     u.Name,
		Quantity: u.Quantity,
		Symbol:   u.Symbol,
	}
}

func mapGoUnitsToUnits(u []goUnits.Unit) []unit {
	var units = []unit{}
	for _, unit := range u {
		units = append(units, mapUnitToDto(unit))
	}
	return units
}

func (u *unitController) listAllUnits(c echo.Context) error {
	return c.JSON(http.StatusOK, mapGoUnitsToUnits(u.unitService.ListAll()))
}

func (u *unitController) listUnitsByQuantity(c echo.Context) error {
	var quantity Quantity
	err := quantity.Scan(c.Param("quantity"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	units, err := u.unitService.ListByQuantity(quantity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mapGoUnitsToUnits(units))
}
