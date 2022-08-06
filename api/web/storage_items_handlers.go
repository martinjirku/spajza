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
	updateFieldRequest struct {
		Value interface{} `json:"value"`
	}
)

var (
	contextKey = "storageItemId"
)

type StorageItemService interface {
	Create(ctx context.Context, storageItem domain.NewStorageItem) (domain.StorageItem, error)
	Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (domain.StorageItem, error)
	UpdateField(ctx context.Context, storageItemI uint, fieldName string, value interface{}) error
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

func (h *storageItemHandler) StorageIdContextProvider(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		storageItemId, err := strconv.ParseUint(c.Param("storageItemId"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Set(contextKey, uint(storageItemId))
		return next(c)
	}
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

func (h *storageItemHandler) updateTitle(c echo.Context) error {
	requestBody := updateFieldRequest{}
	err := c.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id := c.Get(contextKey).(uint)
	err = h.storageItemService.UpdateField(c.Request().Context(), id, "title", requestBody.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusNoContent)

}

func (h *storageItemHandler) consumpt(c echo.Context) error {
	consumptRequest := consumptRequest{}
	err := c.Bind(&consumptRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id := c.Get(contextKey).(uint)
	response, err := h.storageItemService.Consumpt(c.Request().Context(), id, consumptRequest.Amount, consumptRequest.Unit)
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
