package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/domain"
	"github.com/martinjirku/zasobar/repository"
	"github.com/martinjirku/zasobar/usecases"
	middleware "github.com/martinjirku/zasobar/web/middleware"
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

func (h *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	data := UserLoginRequest{}
	json.Unmarshal(body, &data)

	err = h.userService.Login(r.Context(), data.Username, data.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	tokenString, err := h.tokenProvider.GetToken(data.Username, time.Now())
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "WrongJwtToken")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int((config.DefaultConfiguration.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	respondNoContent(w)
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	data := UserRegistrationRequest{}
	err := bindBody(r, &data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	response, err := h.userService.Register(r.Context(), data.Username, data.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, UserRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func (h *UserHandler) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	respondNoContent(w)
}

func (h *UserHandler) aboutMe(w http.ResponseWriter, r *http.Request) {
	userVal := r.Context().Value(middleware.UserKey)
	if userVal == nil {
		respondWithError(w, http.StatusBadRequest, "JwtMalformedSubNotProvided")
		return
	}
	user := userVal.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"]
	if username, ok := sub.(string); ok {
		respondWithJSON(w, http.StatusOK, UserMeResponse{Username: username})
		return
	} else {
		respondWithError(w, http.StatusBadRequest, "JwtMalformedSubNotProvided")
	}
}
