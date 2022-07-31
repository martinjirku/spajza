package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinjirku/zasobar/config"
)

func CreateWebServer(port string) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.DefaultConfiguration.JWTSecret),
		TokenLookup: "cookie:auth",
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == "/api/user/login"
		},
	}))

	// users
	e.POST("/api/user/login", loginHandler)
	e.POST("/api/user/register", registerHandler)
	e.POST("/api/user/logout", logoutHandler)
	e.GET("/api/user/me", aboutMeHandler)

	// units
	e.GET("/api/units", listHandler)
	e.GET("/api/units/:quantity", listUnitsByQuantityHandler)

	// categories
	e.GET("/api/categories", listCategoriesHandler)
	e.POST("/api/categories", saveCategoryHandler)
	e.POST("/api/categories/:id", saveCategoryHandler)
	e.DELETE("/api/categories/:id", deleteCategoryHandler)

	// storage places
	e.POST("/api/storage/places", createStoragePlaceHandler)
	e.GET("/api/storage/places", listStoragePlaceHandler)
	e.POST("/api/storage/places/:storagePlaceId", updateStoragePlaceHandler)
	e.GET("/api/storage/places/:storagePlaceId", getStoragePlaceHandler)
	e.DELETE("/api/storage/places/:storagePlaceId", deleteStoragePlaceHandler)

	// storage items
	e.POST("/api/storage/items", createStorageItemHandler)
	return e, nil
}

func StartWebServer(e *echo.Echo) error {
	return e.Start(config.DefaultConfiguration.Domain + ":" + config.DefaultConfiguration.Port)
}
