package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
)

type storageItemHandler struct{}

func createStorageItemHandler() *storageItemHandler {
	return &storageItemHandler{}
}

func (h *storageItemHandler) createStorageItem(c echo.Context) error {
	storageService := repository.NewStorageService(repository.SqlDb)
	requestBody := domain.NewStorageItemRequest{}
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
