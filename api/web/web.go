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

	return e, nil
}

func StartWebServer(e *echo.Echo) error {
	return e.Start(config.DefaultConfiguration.Domain + ":" + config.DefaultConfiguration.Port)
}
