package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/db"
	"github.com/martinjirku/zasobar/storage"
)

type (
	submitStoragePlaceRequestDto struct {
		Title string `json:"title"`
		Code  string `json:"code"`
	}
	storagePlaceResponseDto struct {
		StoragePlaceId uint   `json:"storagePlaceId"`
		Title          string `json:"title,omitempty"`
		Code           string `json:"code"`
	}
	updateStoragePlaceRequestDto struct {
		Title string
	}
)

func createStoragePlaceHandler(ctx echo.Context) error {
	storagePlacesService := storage.NewStoragePlacesService(db.SqlDb)
	var storagePlace = storage.StoragePlace{}
	err := ctx.Bind(&storagePlace)
	if err != nil {
		return echo.ErrBadRequest
	}
	storagePlace, err = storagePlacesService.Create(ctx.Request().Context(), storagePlace)
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

func updateStoragePlaceHandler(ctx echo.Context) error {
	storagePlacesService := storage.NewStoragePlacesService(db.SqlDb)
	storagePlace := storage.StoragePlace{}
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
	storagePlace, err = storagePlacesService.Update(ctx.Request().Context(), storagePlace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func getStoragePlaceHandler(ctx echo.Context) error {
	storagePlacesService := storage.NewStoragePlacesService(db.SqlDb)
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	storagePlace, err := storagePlacesService.Get(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func deleteStoragePlaceHandler(ctx echo.Context) error {
	storagePlacesService := storage.NewStoragePlacesService(db.SqlDb)
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = storagePlacesService.Delete(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func listStoragePlaceHandler(ctx echo.Context) error {
	storagePlacesService := storage.NewStoragePlacesService(db.SqlDb)
	storagePlaces, err := storagePlacesService.List(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := make([]storagePlaceResponseDto, 0)
	for _, p := range storagePlaces {
		response = append(response, mapStoragePlaceResponseToDto(p))
	}
	return ctx.JSON(http.StatusOK, response)
}

func mapStoragePlaceResponseToDto(storagePlace storage.StoragePlace) storagePlaceResponseDto {
	return storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code}
}
