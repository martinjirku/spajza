package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (h *categoryHandler) listCategories(c echo.Context) error {
	response := []listAllResponse{}
	categories, err := h.categoryService.ListAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for _, c := range categories {
		response = append(response, listAllResponse{
			Id:          c.ID,
			Title:       c.Title,
			Path:        c.Path,
			DefaultUnit: c.DefaultUnit,
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) saveCategory(c echo.Context) error {
	providedCategory := categoryItemDto{}
	if err := c.Bind(&providedCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	idStr := c.Param("id")
	var category = mapCategoryItemToCategory(providedCategory)
	if idStr == "" {
		response, err := h.categoryService.CreateItem(c.Request().Context(), category)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, mapCategoryToCategoryItem(response))
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	category.ID = uint(id)
	response, err := h.categoryService.UpdateItem(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mapCategoryToCategoryItem(response))
}

func (h *categoryHandler) deleteCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.categoryService.DeleteItem(c.Request().Context(), uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}
