package web

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/db"
	"github.com/martinjirku/zasobar/users"
)

type (
	UserRegistrationRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserRegistrationResponse struct {
		Username string `json:"username"`
		Id       int    `json:"password"`
	}
	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserMeResponse struct {
		Username string `json:"username"`
	}
)

func createUserService() users.UserService {
	return users.NewUserService(db.SqlDb)
}

func loginHandler(c echo.Context) error {
	userService := createUserService()
	data := UserLoginRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err := userService.Login(c.Request().Context(), data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	tokenProvider := users.NewTokenProvider(config.DefaultConfiguration.JWTSecret, config.DefaultConfiguration.JWTValidity, config.DefaultConfiguration.JWTIssuer)
	tokenString, err := tokenProvider.GetToken(data.Username, time.Now())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "WrongJwtToken")
	}
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int((config.DefaultConfiguration.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func registerHandler(c echo.Context) error {
	userService := createUserService()
	data := &UserRegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "Bad request"})
	}

	response, err := userService.Register(c.Request().Context(), data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, UserRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func logoutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func aboutMeHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"]
	if username, ok := sub.(string); ok {
		c.JSON(http.StatusOK, UserMeResponse{Username: username})
		return nil
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "JwtMalformedSubNotProvided"})
	}
}
