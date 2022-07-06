package users

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/martinjirku/zasobar/config"
)

type (
	userRegistrationRequest struct {
		Username string
		Password string
	}
	userRegistrationResponse struct {
		Username string
		Id       int
	}
	userLoginRequest struct {
		Username string
		Password string
	}
	userController struct {
		config      config.Configuration
		userService UserService
	}
	UserMeResponse struct {
		Username string `json:"username"`
	}
)

func newUserController(userRepository UserService, config config.Configuration) userController {
	return userController{userService: userRepository, config: config}
}

func (h *userController) register(c echo.Context) error {
	data := &userRegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"Message": "Bad request"})
	}

	response, err := h.userService.Register(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, userRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func (h *userController) login(c echo.Context) error {
	data := userLoginRequest{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	err := h.userService.Login(data.Username, data.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	tokenProvider := NewTokenProvider(h.config.JWTSecret, h.config.JWTValidity, h.config.JWTIssuer)
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

func (h *userController) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	return c.NoContent(http.StatusNoContent)
}

func (h *userController) AboutMe(c echo.Context) error {
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
