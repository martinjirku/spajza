package main

import (
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/controller"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/services"
	"github.com/martinjirku/zasobar/storage"
)

func main() {
	e := echo.New()
	repository := initRepository()

	userController := controller.NewUserController(&repository.User, config.DefaultConfiguration)

	e.POST("/api/user/register", userController.Register)
	e.POST("/api/user/login", userController.Login)
	e.POST("/api/user/logout", userController.Logout)
	e.Logger.Fatal(e.Start(config.DefaultConfiguration.Domain + ":" + config.DefaultConfiguration.Port))
}

func initRepository() services.RepositoryService {
	db := storage.NewDB()
	storage.AutoMigrate(db)
	return services.NewRepositoryService(db)
}
