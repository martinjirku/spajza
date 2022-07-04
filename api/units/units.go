package units

import (
	"github.com/labstack/echo/v4"
)

type UnitsApp struct {
	controller UnitController
}

func NewUnitApp() UnitsApp {
	unitController := NewUnitController()
	users := UnitsApp{controller: unitController}
	return users
}

func (u *UnitsApp) SetupRouter(e *echo.Echo) {
	e.GET("/api/units", u.controller.ListAllUnits)
	e.GET("/api/units/:quantity", u.controller.ListUnitsByQuantity)
}
