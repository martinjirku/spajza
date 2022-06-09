package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/services"
)

type (
	UserRegistrationRequest struct {
		Username string
		Password string
	}
	UserLoginRequest struct {
		Username string
		Password string
	}
	UserController struct {
		config      *config.Configuration
		userService *services.UserService
	}
)

func NewUserController(userRepository *services.UserService, config *config.Configuration) UserController {
	return UserController{userService: userRepository, config: config}
}

func (h *UserController) Register(c echo.Context) error {
	data := UserRegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	response, err := h.userService.Register(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Conflict")
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserController) Login(c echo.Context) error {
	data := UserLoginRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err := h.userService.Login(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "WrongCredentials")
	}
	tokenProvider := services.NewTokenProvider(h.config.JWTSecret, h.config.JWTValidity, h.config.JWTIssuer)
	tokenString, err := tokenProvider.GetToken(data.Username, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "WrongJwtToken")
	}
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    *tokenString,
		MaxAge:   int((h.config.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}
