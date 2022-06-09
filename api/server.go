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
	e.GET("/api/user/login", userController.Login)
	e.Logger.Fatal(e.Start(":1323"))
}

func initRepository() services.RepositoryService {
	db := storage.NewDB()
	storage.AutoMigrate(db)
	return services.NewRepositoryService(db)
}
