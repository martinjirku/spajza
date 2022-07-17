package storage

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoragePlacesController struct {
	storagePlacesService StoragePlacesService
}

type (
	SubmitStoragePlaceRequestDto struct {
		Title string `json:"title"`
		Code  string `json:"code"`
	}
	StoragePlaceResponseDto struct {
		StoragePlaceId uint   `json:"storagePlaceId"`
		Title          string `json:"title,omitempty"`
		Code           string `json:"code"`
	}
	UpdateStoragePlaceRequestDto struct {
		Title string
	}
)

func NewStoragePlacesController(storagePlacesService StoragePlacesService) StoragePlacesController {
	return StoragePlacesController{storagePlacesService}
}

func (s *StoragePlacesController) create(ctx echo.Context) error {
	var storagePlace = StoragePlace{}
	err := ctx.Bind(&storagePlace)
	if err != nil {
		return echo.ErrBadRequest
	}
	storagePlace, err = s.storagePlacesService.Create(ctx.Request().Context(), storagePlace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := StoragePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (s *StoragePlacesController) update(ctx echo.Context) error {
	var storagePlace = StoragePlace{}
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
	storagePlace, err = s.storagePlacesService.Update(ctx.Request().Context(), storagePlace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func (s *StoragePlacesController) get(ctx echo.Context) error {
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	storagePlace, err := s.storagePlacesService.Get(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	return ctx.JSON(http.StatusOK, response)
}

func (s *StoragePlacesController) delete(ctx echo.Context) error {
	storagePlaceIdAsStr := ctx.Param("storagePlaceId")
	storagePlaceId, err := strconv.ParseUint(storagePlaceIdAsStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = s.storagePlacesService.Delete(ctx.Request().Context(), uint(storagePlaceId))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *StoragePlacesController) list(ctx echo.Context) error {
	storagePlaces, err := s.storagePlacesService.List(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	response := make([]StoragePlaceResponseDto, 0)
	for _, p := range storagePlaces {
		response = append(response, mapStoragePlaceResponseToDto(p))
	}
	return ctx.JSON(http.StatusOK, response)
}

func mapStoragePlaceResponseToDto(storagePlace StoragePlace) StoragePlaceResponseDto {
	return StoragePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code}
}
