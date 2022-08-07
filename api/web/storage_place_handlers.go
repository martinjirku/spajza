package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
	"github.com/martinjirku/zasobar/usecases"
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
	storagePlaceRepository := repository.NewStoragePlaceRepository(repository.SqlDb)
	storagePlaceService := usecases.NewStoragePlaceService(storagePlaceRepository)
	return &storagePlaceHandler{storagePlaceService}
}

func (h *storagePlaceHandler) createStoragePlace(w http.ResponseWriter, r *http.Request) {
	storagePlace := domain.StoragePlace{}
	err := bindBody(r, &storagePlace)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	storagePlace, err = h.storagePlaceService.Create(r.Context(), storagePlace)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code,
	}
	respondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) updateStoragePlace(w http.ResponseWriter, r *http.Request) {
	storagePlace := domain.StoragePlace{}
	err := bindBody(r, &storagePlace)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	storagePlace.StoragePlaceId = uint(id)
	storagePlace, err = h.storagePlaceService.Update(r.Context(), storagePlace)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	respondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) getStoragePlace(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	storagePlace, err := h.storagePlaceService.Get(r.Context(), uint(id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	respondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) deleteStoragePlace(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	err = h.storagePlaceService.Delete(r.Context(), uint(id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	respondNoContent(w)
}

func (h *storagePlaceHandler) listStoragePlace(w http.ResponseWriter, r *http.Request) {
	storagePlaces, err := h.storagePlaceService.List(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := make([]storagePlaceResponseDto, 0)
	for _, p := range storagePlaces {
		response = append(response, mapStoragePlaceResponseToDto(p))
	}
	respondWithJSON(w, http.StatusOK, response)
}

func mapStoragePlaceResponseToDto(storagePlace domain.StoragePlace) storagePlaceResponseDto {
	return storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code}
}
