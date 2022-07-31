package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
)

type (
	storagePlaceResponseDto struct {
		StoragePlaceId uint   `json:"storagePlaceId"`
		Title          string `json:"title,omitempty"`
		Code           string `json:"code"`
	}
)

type StoragePlaceService interface {
	Create(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error)
	Get(ctx context.Context, storagePlaceId uint) (domain.StoragePlace, error)
	List(ctx context.Context) ([]domain.StoragePlace, error)
	Update(ctx context.Context, storagePlace domain.StoragePlace) (domain.StoragePlace, error)
	Delete(ctx context.Context, storagePlaceId uint) error
}

type storagePlaceHandler struct {
	storagePlaceService StoragePlaceService
}

func createStoragePlaceHandler() *storagePlaceHandler {
	storagePlaceService := repository.NewStoragePlaceRepository(repository.SqlDb)
	return &storagePlaceHandler{storagePlaceService}
}

func (h *storagePlaceHandler) createStoragePlace(ctx echo.Context) error {
	var storagePlace = domain.StoragePlace{}
	err := ctx.Bind(&storagePlace)
	if err != nil {
		return echo.ErrBadRequest
	}
	storagePlace, err = h.storagePlaceService.Create(ctx.Request().Context(), storagePlace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *storagePlaceHandler) updateStoragePlace(ctx echo.Context) error {
	storagePlace := domain.StoragePlace{}
	err := ctx.Bind(&storagePlace)
	if err != nil {
		return echo.ErrBadRequest
	}
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	storagePlace.StoragePlaceId = uint(storagePlaceId)
	storagePlace, err = h.storagePlaceService.Update(ctx.Request().Context(), storagePlace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func (h *storagePlaceHandler) getStoragePlace(ctx echo.Context) error {
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	storagePlace, err := h.storagePlaceService.Get(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func (h *storagePlaceHandler) deleteStoragePlace(ctx echo.Context) error {
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.storagePlaceService.Delete(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *storagePlaceHandler) listStoragePlace(ctx echo.Context) error {
	storagePlaces, err := h.storagePlaceService.List(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := make([]storagePlaceResponseDto, 0)
	for _, p := range storagePlaces {
		response = append(response, mapStoragePlaceResponseToDto(p))
	}
	return ctx.JSON(http.StatusOK, response)
}

func mapStoragePlaceResponseToDto(storagePlace domain.StoragePlace) storagePlaceResponseDto {
	return storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code}
}
