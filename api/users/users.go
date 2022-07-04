package users

import (
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"gorm.io/gorm"
)

type Users struct {
	controller UserController
}

func NewUserApp(db *gorm.DB) Users {
	userController := NewUserController(NewUserService(db), config.DefaultConfiguration)
	users := Users{controller: userController}
	return users
}

func (u *Users) SetupRouter(e *echo.Echo) {

	e.POST("/api/user/login", u.controller.Login)
	e.POST("/api/user/register", u.controller.Register)
	e.POST("/api/user/logout", u.controller.Logout)
	e.GET("/api/user/me", u.controller.AboutMe)
}
