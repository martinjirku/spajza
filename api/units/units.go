package units

import (
	"github.com/labstack/echo/v4"
)

func StartApp(e *echo.Echo) {
	unitController := newUnitController()
	e.GET("/api/units", unitController.listAllUnits)
	e.GET("/api/units/:quantity", unitController.listUnitsByQuantity)
}
