package categories

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartApp(db *gorm.DB, e *echo.Echo) {
	service := NewCategoryService(db)
	controller := NewController(service)
	e.GET("/api/categories", controller.ListAll)
	e.POST("/api/categories", controller.SaveCategory)
	e.POST("/api/categories/:id", controller.SaveCategory)
	e.DELETE("/api/categories/:id", controller.DeleteCategory)
}
