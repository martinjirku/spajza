package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/services"
)

type (
	UserRegistrationRequest struct {
		Username string
		Password string
	}
	UserRegistrationResponse struct {
		Username string
		Id       int
	}
	UserLoginRequest struct {
		Username string
		Password string
	}
	UserController struct {
		config      config.Configuration
		userService services.UserService
	}
	UserMeResponse struct {
		Username string `json:"username"`
	}
)

func NewUserController(userRepository services.UserService, config config.Configuration) UserController {
	return UserController{userService: userRepository, config: config}
}

func (h *UserController) Register(c echo.Context) error {
	data := &UserRegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "Bad request"})
	}

	response, err := h.userService.Register(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, UserRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func (h *UserController) Login(c echo.Context) error {
	data := UserLoginRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err := h.userService.Login(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	tokenProvider := services.NewTokenProvider(h.config.JWTSecret, h.config.JWTValidity, h.config.JWTIssuer)
	tokenString, err := tokenProvider.GetToken(data.Username, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "WrongJwtToken")
	}
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int((h.config.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func (h *UserController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func (h *UserController) AboutMe(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"]
	if username, ok := sub.(string); ok {
		return c.JSON(http.StatusOK, UserMeResponse{Username: username})
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "JwtMalformedSubNotProvided"})
	}

}
