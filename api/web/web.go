package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinjirku/zasobar/config"
)

func CreateWebServer(port string) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_unix}",` +
			`"request":"${method} ${uri}",` +
			`"status":${status},"error":"${error}"}` + "\n",
		// Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		// 	`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		// 	`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		// 	`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.DefaultConfiguration.JWTSecret),
		TokenLookup: "cookie:auth",
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == "/api/user/login"
		},
	}))

	// users
	user := createUserHandler()
	e.POST("/api/user/login", user.login)
	e.POST("/api/user/register", user.register)
	e.POST("/api/user/logout", user.logout)
	e.GET("/api/user/me", user.aboutMe)

	// units
	units := createUnitHandler()
	e.GET("/api/units", units.list)
	e.GET("/api/units/:quantity", units.listUnitsByQuantity)

	// categories
	categories := createCategoryHandler()
	e.GET("/api/categories", categories.listCategories)
	e.POST("/api/categories", categories.saveCategory)
	e.POST("/api/categories/:id", categories.saveCategory)
	e.DELETE("/api/categories/:id", categories.deleteCategory)

	// storage places
	storagePlaceHandler := createStoragePlaceHandler()
	e.POST("/api/storage/places", storagePlaceHandler.createStoragePlace)
	e.GET("/api/storage/places", storagePlaceHandler.listStoragePlace)
	e.POST("/api/storage/places/:storagePlaceId", storagePlaceHandler.updateStoragePlace)
	e.GET("/api/storage/places/:storagePlaceId", storagePlaceHandler.getStoragePlace)
	e.DELETE("/api/storage/places/:storagePlaceId", storagePlaceHandler.deleteStoragePlace)

	// storage items
	storageItemHandler := createStorageItemHandler()
	e.GET("/api/storage/items", storageItemHandler.list)
	e.POST("/api/storage/items", storageItemHandler.createStorageItem)
	storageItemGroup := e.Group("/api/storage/items/:storageItemId")
	storageItemGroup.Use(storageItemHandler.StorageIdContextProvider)
	storageItemGroup.POST("/title", storageItemHandler.updateTitle)
	storageItemGroup.POST("/consumpt", storageItemHandler.consumpt)

	return e, nil
}

func StartWebServer(e *echo.Echo) error {
	return e.Start(config.DefaultConfiguration.Domain + ":" + config.DefaultConfiguration.Port)
}
