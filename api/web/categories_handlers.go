package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/categories"
	"github.com/martinjirku/zasobar/db"
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

func mapCategoryItemToCategory(c categoryItemDto) categories.Category {
	return categories.Category{
		ID:          c.Id,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

func mapCategoryToCategoryItem(c categories.Category) categoryItemDto {
	return categoryItemDto{
		Id:          c.ID,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

func listCategoriesHandler(c echo.Context) error {
	categoryService := categories.NewCategoryService(db.SqlDb)
	var response = []listAllResponse{}
	var categories, err = categoryService.ListAll(c.Request().Context())
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

func saveCategoryHandler(c echo.Context) error {
	categoryService := categories.NewCategoryService(db.SqlDb)
	providedCategory := categoryItemDto{}
	if err := c.Bind(&providedCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	idStr := c.Param("id")
	var category = mapCategoryItemToCategory(providedCategory)
	if idStr == "" {
		response, err := categoryService.CreateItem(c.Request().Context(), category)
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
	response, err := categoryService.UpdateItem(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mapCategoryToCategoryItem(response))

}

func deleteCategoryHandler(c echo.Context) error {
	categoryService := categories.NewCategoryService(db.SqlDb)
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var category = categories.Category{}
	category.ID = uint(id)
	err = categoryService.DeleteItem(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}
