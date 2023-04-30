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

type StoragePlaceService interface {
	Create(storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Get(storagePlaceId uint) (entity.StoragePlace, error)
	List() ([]entity.StoragePlace, error)
	Update(storagePlace entity.StoragePlace) (entity.StoragePlace, error)
	Delete(storagePlaceId uint) error
}

type storagePlaceHandler struct {
	db *sql.DB
}

func CreateStoragePlaceHandler(db *sql.DB) *storagePlaceHandler {
	return &storagePlaceHandler{db}
}

func (h *storagePlaceHandler) getUsecase(ctx context.Context) *usecase.StoragePlaceUsecase {
	storagePlaceRepository := repository.NewStoragePlaceRepository(ctx, h.db)
	return usecase.NewStoragePlaceUsecase(storagePlaceRepository)
}

func (h *storagePlaceHandler) CreateStoragePlace(w http.ResponseWriter, r *http.Request) {
	storagePlace := entity.StoragePlace{}
	err := web.BindBody(r, &storagePlace)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	usecase := h.getUsecase(r.Context())
	storagePlace, err = usecase.Create(storagePlace)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code,
	}
	web.RespondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) UpdateStoragePlace(w http.ResponseWriter, r *http.Request) {
	storagePlace := entity.StoragePlace{}
	err := web.BindBody(r, &storagePlace)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	storagePlace.StoragePlaceId = int32(id)
	usecase := h.getUsecase(r.Context())
	storagePlace, err = usecase.Update(storagePlace)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	web.RespondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) GetStoragePlace(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	usecase := h.getUsecase(r.Context())
	storagePlace, err := usecase.Get(int32(id))
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := mapStoragePlaceResponseToDto(storagePlace)
	web.RespondWithJSON(w, http.StatusOK, response)
}

func (h *storagePlaceHandler) DeleteStoragePlace(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	usecase := h.getUsecase(r.Context())
	err = usecase.Delete(int32(id))
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	web.RespondNoContent(w)
}

func (h *storagePlaceHandler) ListStoragePlace(w http.ResponseWriter, r *http.Request) {
	usecase := h.getUsecase(r.Context())
	storagePlaces, err := usecase.List()
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	response := make([]storagePlaceResponseDto, 0)
	for _, p := range storagePlaces {
		response = append(response, mapStoragePlaceResponseToDto(p))
	}
	web.RespondWithJSON(w, http.StatusOK, response)
}

func mapStoragePlaceResponseToDto(storagePlace entity.StoragePlace) storagePlaceResponseDto {
	return storagePlaceResponseDto{
		StoragePlaceId: storagePlace.StoragePlaceId,
		Title:          storagePlace.Title,
		Code:           storagePlace.Code}
}
