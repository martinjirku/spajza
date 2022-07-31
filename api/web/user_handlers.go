package web

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
	"github.com/martinjirku/zasobar/usecases"
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

type UserService interface {
	ListAll() ([]*domain.User, error)
	Register(ctx context.Context, email string, password string) (domain.User, error)
	Login(ctx context.Context, email string, password string) error
}

type TokenProvider interface {
	GetToken(userName string, currentTime time.Time) (string, error)
}

type UserHandler struct {
	userService   UserService
	tokenProvider TokenProvider
}

func createUserHandler() *UserHandler {
	userService := repository.NewUserRepository(repository.SqlDb)
	tokenProvider := usecases.NewTokenProvider(config.DefaultConfiguration.JWTSecret, config.DefaultConfiguration.JWTValidity, config.DefaultConfiguration.JWTIssuer)
	return &UserHandler{userService, tokenProvider}
}

func (h *UserHandler) login(c echo.Context) error {
	data := UserLoginRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err := h.userService.Login(c.Request().Context(), data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	tokenString, err := h.tokenProvider.GetToken(data.Username, time.Now())
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

func (h *UserHandler) register(c echo.Context) error {
	data := &UserRegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "Bad request"})
	}

	response, err := h.userService.Register(c.Request().Context(), data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, UserRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func (h *UserHandler) logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) aboutMe(c echo.Context) error {
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
