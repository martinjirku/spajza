package units

import (
	"net/http"

	goUnits "github.com/bcicen/go-units"
	"github.com/labstack/echo/v4"
)

type (
	UnitController struct {
		unitService UnitService
	}

	Unit struct {
		Name     string
		Symbol   string
		Type     string
		Quantity string
	}
)

func NewUnitController() UnitController {
	return UnitController{unitService: NewUnitService()}
}

func mapUnitToDto(u goUnits.Unit) Unit {
	return Unit{
		Name:     u.Name,
		Quantity: u.Quantity,
		Symbol:   u.Symbol,
	}
}

func mapGoUnitsToUnits(u []goUnits.Unit) []Unit {
	var units = []Unit{}
	for _, unit := range u {
		units = append(units, mapUnitToDto(unit))
	}
	return units
}

func (u *UnitController) ListAllUnits(c echo.Context) error {
	return c.JSON(http.StatusOK, mapGoUnitsToUnits(u.unitService.ListAll()))
}

func (u *UnitController) ListUnitsByQuantity(c echo.Context) error {
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
