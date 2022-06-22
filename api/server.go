package main

import (
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinjirku/zasobar/services"
	"github.com/martinjirku/zasobar/storage"
)

func main() {
	e := echo.New()
	repository := initRepository()

	userController := controller.NewUserController(repository.User, config.DefaultConfiguration)
	unitController := controller.NewUnitController()
	e.Use(middleware.Logger())
	e.Logger.Info(config.DefaultConfiguration.JWTSecret)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.DefaultConfiguration.JWTSecret),
		TokenLookup: "cookie:auth",
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == "/api/user/login"
		},
	}))

	e.POST("/api/user/login", userController.Login)
	e.POST("/api/user/register", userController.Register)
	e.POST("/api/user/logout", userController.Logout)
	e.GET("/api/user/me", userController.AboutMe)
	e.GET("/api/units/possible-units", unitController.GetPossibleUnits)
	e.Logger.Fatal(e.Start(config.DefaultConfiguration.Domain + ":" + config.DefaultConfiguration.Port))
}

func initRepository() services.RepositoryService {
	db := storage.NewDB()
	return services.NewRepositoryService(db)
}
