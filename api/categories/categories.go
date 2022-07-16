package categories

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func StartApp(db *sql.DB, e *echo.Echo) {
	service := NewCategoryService(db)
	controller := NewController(service)
	e.GET("/api/categories", controller.ListAll)
	e.POST("/api/categories", controller.SaveCategory)
	e.POST("/api/categories/:id", controller.SaveCategory)
	e.DELETE("/api/categories/:id", controller.DeleteCategory)
}
