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

type CategoryService interface {
	ListAll() ([]entity.Category, error)
	CreateItem(c entity.Category) (entity.Category, error)
	UpdateItem(c entity.Category) (entity.Category, error)
	DeleteItem(id uint) error
}

type categoryHandler struct {
	db *sql.DB
}

func CreateCategoryHandler(db *sql.DB) *categoryHandler {
	return &categoryHandler{db}
}

func (h *categoryHandler) getUsecase(ctx context.Context) CategoryService {
	categoryRepository := repository.NewCategoryService(ctx, h.db)
	return usecase.CreateCategoryService(categoryRepository)
}

func (h *categoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	response := []listAllResponse{}

	usecase := h.getUsecase(r.Context())
	categories, err := usecase.ListAll()
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	for _, c := range categories {
		response = append(response, listAllResponse{
			Id:          c.ID,
			Title:       c.Title,
			Path:        c.Path,
			DefaultUnit: c.DefaultUnit,
		})
	}
	web.RespondWithJSON(w, http.StatusOK, response)
}

func (h *categoryHandler) SaveCategory(w http.ResponseWriter, r *http.Request) {
	usecase := h.getUsecase(r.Context())
	providedCategory := categoryItemDto{}
	if err := web.BindBody(r, &providedCategory); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	category := mapCategoryItemToCategory(providedCategory)
	if idStr == "" {
		response, err := usecase.CreateItem(category)
		if err != nil {
			web.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		web.RespondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	category.ID = uint(id)
	response, err := usecase.UpdateItem(category)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
}

func (h *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	usecase := h.getUsecase(r.Context())
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = usecase.DeleteItem(uint(id))
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondNoContent(w)
}
