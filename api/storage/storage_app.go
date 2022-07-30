package storage

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func StartApp(db *sql.DB, e *echo.Echo) {
	storagePlacesService := NewStoragePlacesService(db)
	storagePlacesController := NewStoragePlacesController(storagePlacesService)
	storageService := NewStorageService(db)
	storageController := newStorageController(storageService)

	e.POST("/api/storage/items", storageController.create)
	e.POST("/api/storage/places", storagePlacesController.create)
	e.GET("/api/storage/places", storagePlacesController.list)
	e.POST("/api/storage/places/:storagePlaceId", storagePlacesController.update)
	e.GET("/api/storage/places/:storagePlaceId", storagePlacesController.get)
	e.DELETE("/api/storage/places/:storagePlaceId", storagePlacesController.delete)
}
