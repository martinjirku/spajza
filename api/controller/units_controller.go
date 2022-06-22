package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/services"
)

type (
	UnitController struct {
	}
)

func NewUnitController() UnitController {
	return UnitController{}
}

func (u *UnitController) GetPossibleUnits(c echo.Context) error {
	return c.JSON(http.StatusOK, services.PossibleUnits())
}
