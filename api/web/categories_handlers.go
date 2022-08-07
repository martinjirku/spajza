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
	categoryItemDto struct {
		Id          uint   `json:"id"`
		Title       string `json:"title"`
		Path        string `json:"path"`
		DefaultUnit string `json:"defaultUnit"`
	}
	listAllResponse categoryItemDto
)

type CategoryService interface {
	ListAll(ctx context.Context) ([]domain.Category, error)
	CreateItem(ctx context.Context, c domain.Category) (domain.Category, error)
	UpdateItem(ctx context.Context, c domain.Category) (domain.Category, error)
	DeleteItem(ctx context.Context, id uint) error
}

type categoryHandler struct {
	categoryService CategoryService
}

func createCategoryHandler() *categoryHandler {
	categoryRepository := repository.NewCategoryService(repository.SqlDb)
	categoryService := usecases.CreateCategoryService(categoryRepository)
	return &categoryHandler{categoryService}
}

func mapCategoryItemToCategory(c categoryItemDto) domain.Category {
	return domain.Category{
		ID:          c.Id,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

func mapCategoryToCategoryItem(c domain.Category) categoryItemDto {
	return categoryItemDto{
		Id:          c.ID,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

func (h *categoryHandler) listCategories(w http.ResponseWriter, r *http.Request) {
	response := []listAllResponse{}

	categories, err := h.categoryService.ListAll(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
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
	respondWithJSON(w, http.StatusOK, response)
}

func (h *categoryHandler) saveCategory(w http.ResponseWriter, r *http.Request) {
	providedCategory := categoryItemDto{}
	if err := bindBody(r, &providedCategory); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := chi.URLParam(r, "id")
	category := mapCategoryItemToCategory(providedCategory)
	if idStr == "" {
		response, err := h.categoryService.CreateItem(r.Context(), category)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	category.ID = uint(id)
	response, err := h.categoryService.UpdateItem(r.Context(), category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, mapCategoryToCategoryItem(response))
}

func (h *categoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.categoryService.DeleteItem(r.Context(), uint(id))
	if err != nil {

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondNoContent(w)
}
