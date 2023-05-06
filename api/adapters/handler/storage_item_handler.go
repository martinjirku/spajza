package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/martinjirku/zasobar/entity"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
)

type StorageItemUsecaseProvider func(ctx context.Context) *usecase.StorageItemUsecase

type usecaseHandler struct {
	provideUsecase StorageItemUsecaseProvider
}

func CreateStorageItemHandler(provider StorageItemUsecaseProvider) usecaseHandler {
	return usecaseHandler{provider}
}

func (h *usecaseHandler) CreateStorageItem(w http.ResponseWriter, r *http.Request) {
	requestBody := NewStorageItem{}
	err := web.BindBody(r, &requestBody)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	usecase := h.provideUsecase(r.Context())
	item := entity.StorageItem{
		StorageItemID:  requestBody.StoragePlaceId,
		Title:          requestBody.Title,
		CategoryID:     requestBody.CategoryId,
		StoragePlaceID: requestBody.StoragePlaceId,
		ExpirationDate: requestBody.ExpirationDate,
		Ean:            requestBody.Ean,
	}
	item.Init()
	item.SetBaselineQuantity(entity.Quantity{Value: requestBody.Amount, Unit: entity.UnitName(requestBody.Unit)})
	item, err = usecase.Create(item)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	response := mapEntityToStorageItem(item)
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *usecaseHandler) UpdateField(w http.ResponseWriter, r *http.Request) {
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
	usecase := h.provideUsecase(r.Context())
	_, err = usecase.UpdateField(int32(id), fieldName, requestBody.Value)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondNoContent(w)
}

func (h *usecaseHandler) Consumpt(w http.ResponseWriter, r *http.Request) {
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
	usecase := h.provideUsecase(r.Context())
	response, err := usecase.Consumpt(int32(id), requestBody.Amount, entity.UnitName(requestBody.Unit))
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusAccepted, response)
}

func (h *usecaseHandler) List(w http.ResponseWriter, r *http.Request) {
	usecase := h.provideUsecase(r.Context())
	items, err := usecase.List()
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	result := make([]StorageItem, len(items))
	for i, si := range items {
		result[i] = mapEntityToStorageItem(si)
	}
	web.RespondWithJSON(w, http.StatusOK, listResponse{result})
}
