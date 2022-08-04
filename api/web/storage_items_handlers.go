package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
	"github.com/martinjirku/zasobar/usecases"
)

type (
	consumptRequest struct {
		Amount float64 `json:"amount"`
		Unit   string  `json:"unit"`
	}
	listResponse struct {
		Items []domain.StorageItem `json:"items"`
	}
)

type StorageItemService interface {
	Create(ctx context.Context, storageItem domain.NewStorageItem) (domain.StorageItem, error)
	Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (domain.StorageItem, error)
	List(ctx context.Context) ([]domain.StorageItem, error)
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
	requestBody := domain.NewStorageItem{}
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

func (h *storageItemHandler) consumpt(c echo.Context) error {
	consumptRequest := consumptRequest{}
	err := c.Bind(&consumptRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	storageItemId, err := strconv.ParseUint(c.Param("storageItemId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response, err := h.storageItemService.Consumpt(c.Request().Context(), uint(storageItemId), consumptRequest.Amount, consumptRequest.Unit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "response")
	}
	return c.JSON(http.StatusAccepted, response)
}

func (h *storageItemHandler) list(c echo.Context) error {
	result, err := h.storageItemService.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, listResponse{result})
}
