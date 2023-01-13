package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/martinjirku/zasobar/adapters/repository"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/entity"
	web "github.com/martinjirku/zasobar/pkg/web"
	"github.com/martinjirku/zasobar/usecase"
	"google.golang.org/api/idtoken"
)

type UserGateway interface {
	Register(string, password string) (entity.User, error)
	Login(email string, password string) error
}

type TokenProvider interface {
	GetToken(userName string, currentTime time.Time) (string, error)
}

type UserHandler struct {
	db            *sql.DB
	config        config.Configuration
	tokenProvider TokenProvider
}

func CreateUserHandler(db *sql.DB, config config.Configuration) *UserHandler {
	tokenProvider := web.NewTokenProvider(config.JWTSecret, config.JWTValidity, config.JWTIssuer)
	return &UserHandler{db, config, tokenProvider}
}

func (h *UserHandler) getUsecase(ctx context.Context) usecase.UserUsecase {
	userRepository := repository.NewUserRepository(ctx, h.db)
	return usecase.GetUserUsecase(userRepository)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	data := UserLoginRequest{}
	json.Unmarshal(body, &data)
	usecase := h.getUsecase(r.Context())
	err = usecase.Login(data.Username, data.Password)
	if err != nil {
		web.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	tokenString, err := h.tokenProvider.GetToken(data.Username, time.Now())
	if err != nil {
		web.RespondWithError(w, http.StatusUnauthorized, "WrongJwtToken")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int((config.DefaultConfiguration.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	web.RespondNoContent(w)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := UserRegistrationRequest{}
	err := web.BindBody(r, &data)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	usecase := h.getUsecase(r.Context())
	response, err := usecase.Register(data.Username, data.Password)
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, UserRegistrationResponse{Id: int(response.ID), Username: response.Email})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	})
	web.RespondNoContent(w)
}

func (h *UserHandler) AboutMe(w http.ResponseWriter, r *http.Request) {
	userVal := r.Context().Value(web.UserKey)
	if userVal == nil {
		web.RespondWithError(w, http.StatusBadRequest, "JwtMalformedSubNotProvided")
		return
	}
	user := userVal.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"]
	if username, ok := sub.(string); ok {
		web.RespondWithJSON(w, http.StatusOK, UserMeResponse{Username: username})
		return
	} else {
		web.RespondWithError(w, http.StatusBadRequest, "JwtMalformedSubNotProvided")
	}
}

func (h *UserHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	redirect := r.URL.Query().Get("redirect")

	r.ParseForm()
	idToken := r.PostForm.Get("credential")
	clientID := "1089933053808-glhibpnso4vbc38beorao10b30p64d84.apps.googleusercontent.com" //"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"
	payload, err := idtoken.Validate(r.Context(), idToken, clientID)
	if err != nil {
		web.RespondWithError(w, 400, fmt.Sprint(err))
		return
	}
	uc := h.getUsecase(r.Context())
	googleOptions := usecase.GoogleProviderOptions{
		Email:         payload.Claims["email"].(string),
		Name:          payload.Claims["name"].(string),
		GivenName:     payload.Claims["given_name"].(string),
		FamilyName:    payload.Claims["family_name"].(string),
		Picture:       payload.Claims["picture"].(string),
		EmailVerified: payload.Claims["email_verified"].(bool),
	}
	errByLogin := uc.LoginByProvider(usecase.LoginOptions{
		Provider:       entity.AuthProviderGoogle,
		GoogleProvider: googleOptions,
	})

	w.Header().Set("Content-Type", "text/html")

	if errByLogin != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head></head><script>window.location="/prihlasenie/?googleFailed"</script></html>`))
		return
	}
	tokenString, err := h.tokenProvider.GetToken(googleOptions.Email, time.Now())

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head></head><script>window.location="/prihlasenie/?googleFailed"</script></html>`))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int((config.DefaultConfiguration.JWTValidity + 2) * 60),
		HttpOnly: true,
	})
	if redirect == "" {
		redirect = "/"
	}
	w.WriteHeader(http.StatusOK)
	html := fmt.Sprintf(`<html><head></head><script>window.location="%s"</script></html>`, redirect)
	w.Write([]byte(html))
}
