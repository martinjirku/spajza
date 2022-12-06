package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/adapters/repository"
	"github.com/martinjirku/zasobar/entity"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
)

type StorageItemService interface {
	Create(ctx context.Context, storageItem entity.NewStorageItem) (entity.StorageItem, error)
	Consumpt(ctx context.Context, storageItemId uint, amount float64, unit string) (entity.StorageItem, error)
	UpdateField(ctx context.Context, storageItemId uint, fieldName string, value interface{}) error
	List(ctx context.Context) ([]entity.StorageItem, error)
}

type storageItemHandler struct {
	db *sql.DB
}

func CreateStorageItemHandler(db *sql.DB) *storageItemHandler {
	return &storageItemHandler{db}
}

func (h *storageItemHandler) getUsecase(ctx context.Context) *usecase.StorageItemService {
	storageItemRepository := repository.NewStorageItemRepository(ctx, h.db)
	return usecase.NewStorageItemService(&storageItemRepository)
}

func (h *storageItemHandler) CreateStorageItem(w http.ResponseWriter, r *http.Request) {
	requestBody := entity.NewStorageItem{}
	err := web.BindBody(r, &requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	usecase := h.getUsecase(r.Context())
	response, err := usecase.Create(requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) UpdateField(w http.ResponseWriter, r *http.Request) {
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
	usecase := h.getUsecase(r.Context())
	err = usecase.UpdateField(uint(id), fieldName, requestBody.Value)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondNoContent(w)
}

func (h *storageItemHandler) Consumpt(w http.ResponseWriter, r *http.Request) {
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
	usecase := h.getUsecase(r.Context())
	response, err := usecase.Consumpt(uint(id), requestBody.Amount, requestBody.Unit)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *storageItemHandler) List(w http.ResponseWriter, r *http.Request) {
	usecase := h.getUsecase(r.Context())
	result, err := usecase.List()
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, listResponse{result})
}
