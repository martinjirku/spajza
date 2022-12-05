package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/adapters/repository"
	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/infra/db"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecases"
)

type (
	consumptRequest struct {
		Amount float64 `json:"amount"`
		Unit   string  `json:"unit"`
	}
	listResponse struct {
		Items []entity.StorageItem `json:"items"`
	}
	updateFieldRequest struct {
		Value interface{} `json:"value"`
	}
)
type StorageItemService interface {
	Create(ctx context.Context, storageItem entity.NewStorageItem) (entity.StorageItem, error)
	Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (entity.StorageItem, error)
	UpdateField(ctx context.Context, storageItemId uint, fieldName string, value interface{}) error
	List(ctx context.Context) ([]entity.StorageItem, error)
}

type storageItemHandler struct {
	storageItemService StorageItemService
}

func createStorageItemHandler() *storageItemHandler {
	storageItemRepository := repository.NewStorageItemRepository(db.SqlDb)
	storageItemService := usecases.NewStorageItemService(&storageItemRepository)
	return &storageItemHandler{storageItemService}
}

func (h *storageItemHandler) createStorageItem(w http.ResponseWriter, r *http.Request) {
	requestBody := entity.NewStorageItem{}
	err := web.BindBody(r, &requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.storageItemService.Create(r.Context(), requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) updateField(w http.ResponseWriter, r *http.Request) {
	requestBody := updateFieldRequest{}
	err := web.BindBody(r, &requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	fieldName := chi.URLParam(r, "fieldName")
	if fieldName != "title" && fieldName != "storagePlaceId" {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.storageItemService.UpdateField(r.Context(), uint(id), fieldName, requestBody.Value)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondNoContent(w)
}

func (h *storageItemHandler) consumpt(w http.ResponseWriter, r *http.Request) {
	requestBody := consumptRequest{}
	err := web.BindBody(r, &requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.storageItemService.Consumpt(r.Context(), uint(id), requestBody.Amount, requestBody.Unit)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) list(w http.ResponseWriter, r *http.Request) {
	result, err := h.storageItemService.List(r.Context())
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, listResponse{result})
}
