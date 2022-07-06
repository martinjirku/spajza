package users

import (
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"gorm.io/gorm"
)

func StartApp(db *gorm.DB, e *echo.Echo) {
	userController := newUserController(NewUserService(db), config.DefaultConfiguration)

	e.POST("/api/user/login", userController.login)
	e.POST("/api/user/register", userController.register)
	e.POST("/api/user/logout", userController.Logout)
	e.GET("/api/user/me", userController.AboutMe)
}
