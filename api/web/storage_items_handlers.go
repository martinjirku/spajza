package web

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
	"github.com/martinjirku/zasobar/usecases"
)

type StorageItemService interface {
	Create(ctx context.Context, storageItem domain.NewStorageItemRequest) (domain.StorageItemResponse, error)
}

type storageItemHandler struct {
	storageItemService StorageItemService
}

func createStorageItemHandler() *storageItemHandler {
	storageItemRepository := repository.NewStorageItemRepository(repository.SqlDb)
	storageItemService := usecases.NewStorageItemService(&storageItemRepository)
	return &storageItemHandler{storageItemService}
}

func (h *storageItemHandler) createStorageItem(c echo.Context) error {
	requestBody := domain.NewStorageItemRequest{}
	err := c.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response, err := h.storageItemService.Create(c.Request().Context(), requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response)
}
