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

func (h *storageItemHandler) createStorageItem(w http.ResponseWriter, r *http.Request) {
	requestBody := domain.NewStorageItem{}
	err := bindBody(r, &requestBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.storageItemService.Create(r.Context(), requestBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) updateTitle(w http.ResponseWriter, r *http.Request) {
	requestBody := updateFieldRequest{}
	err := bindBody(r, &requestBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	err = h.storageItemService.UpdateField(r.Context(), uint(id), "title", requestBody.Value)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondNoContent(w)
}

func (h *storageItemHandler) consumpt(w http.ResponseWriter, r *http.Request) {
	requestBody := consumptRequest{}
	err := bindBody(r, &requestBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.storageItemService.Consumpt(r.Context(), uint(id), requestBody.Amount, requestBody.Unit)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) list(w http.ResponseWriter, r *http.Request) {
	result, err := h.storageItemService.List(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, listResponse{result})
}
