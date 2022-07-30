package storage

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	ss StorageService
}

func newStorageController(ss StorageService) controller {
	return controller{ss: ss}
}

func (ctrl controller) create(c echo.Context) error {
	var requestBody = NewStorageItemRequest{}
	err := c.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response, err := ctrl.ss.Create(c.Request().Context(), requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response)
}
