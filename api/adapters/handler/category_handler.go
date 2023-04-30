package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
)

type CategoryUsecaseProvider func(ctx context.Context) *usecase.CategoryUsecase

type categoryHandler struct {
	provideUsecase CategoryUsecaseProvider
}

func CreateCategoryHandler(provider CategoryUsecaseProvider) categoryHandler {
	return categoryHandler{provider}
}

func (h *categoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	response := []listAllResponse{}

	usecase := h.provideUsecase(r.Context())
	categories, err := usecase.List()
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	for _, c := range categories {
		response = append(response, listAllResponse{
			Id:          c.ID,
			Title:       c.Title,
			Path:        string(c.Path),
			DefaultUnit: c.DefaultUnit,
		})
	}
	web.RespondWithJSON(w, http.StatusOK, response)
}

func (h *categoryHandler) SaveCategory(w http.ResponseWriter, r *http.Request) {
	usecase := h.provideUsecase(r.Context())
	providedCategory := categoryItemDto{}
	if err := web.BindBody(r, &providedCategory); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	category := mapCategoryItemToCategory(providedCategory)
	if idStr == "" {
		response, err := usecase.Create(category)
		if err != nil {
			web.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		web.RespondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	category.ID = int32(id)
	response, err := usecase.Update(category)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
}

func (h *categoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	usecase := h.provideUsecase(r.Context())
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = usecase.Delete(int32(id))
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondNoContent(w)
}
