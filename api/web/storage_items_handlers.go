package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/db"
	"github.com/martinjirku/zasobar/storage"
)

func createStorageItemHandler(c echo.Context) error {
	storageService := storage.NewStorageService(db.SqlDb)
	var requestBody = storage.NewStorageItemRequest{}
	err := c.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response, err := storageService.Create(c.Request().Context(), requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response)
}
