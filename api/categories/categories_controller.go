package categories

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	cs CategoryService
}

type (
	categoryItemDto struct {
		Id          uint   `json:"id"`
		Title       string `json:"title"`
		Path        string `json:"path"`
		DefaultUnit string `json:"defaultUnit"`
	}
	listAllResponse categoryItemDto
)

func NewController(cs CategoryService) controller {
	return controller{cs: cs}
}

func mapCategoryItemToCategory(c categoryItemDto) Category {
	return Category{
		ID:          c.Id,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}
func mapCategoryToCategoryItem(c Category) categoryItemDto {
	return categoryItemDto{
		Id:          c.ID,
		Title:       c.Title,
		Path:        c.Path,
		DefaultUnit: c.DefaultUnit,
	}
}

func (ctrl *controller) ListAll(c echo.Context) error {
	var response = []listAllResponse{}
	var categories, err = ctrl.cs.ListAll(c.Request().Context())
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

func (ctrl *controller) SaveCategory(c echo.Context) error {
	providedCategory := new(categoryItemDto)
	if err := c.Bind(providedCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	idStr := c.Param("id")
	var category = mapCategoryItemToCategory(*providedCategory)
	if idStr == "" {
		response, err := ctrl.cs.CreateItem(c.Request().Context(), category)
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
	response, err := ctrl.cs.UpdateItem(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, mapCategoryToCategoryItem(response))

}

func (ctrl *controller) DeleteCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var category = Category{}
	category.ID = uint(id)
	err = ctrl.cs.DeleteItem(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, "")
}
