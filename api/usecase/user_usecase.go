package usecase

import (
	"strings"

	"github.com/martinjirku/zasobar/entity"
)

type UserUsecase struct {
	gateway UserGateway
}

type UserGateway interface {
	Login(email string, password string) error
	Register(email string, password string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
}

func GetUserUsecase(registrator UserGateway) UserUsecase {
	return UserUsecase{registrator}
}

func (u *UserUsecase) Register(email, password string) (*entity.User, error) {
	return u.gateway.Register(email, password)
}

func (u *UserUsecase) Login(email, password string) error {
	return u.gateway.Login(email, password)
}

func (u *UserUsecase) LoginByGoogle(options GoogleProviderOptions) error {
	userEmail := strings.ToLower(options.Email)
	oldUser, err := u.gateway.GetByEmail(userEmail)
	if err != nil {
		user := entity.User{
			Email:         userEmail,
			Name:          options.Name,
			GivenName:     options.GivenName,
			FamilyName:    options.FamilyName,
			Picture:       options.Picture,
			EmailVerified: options.EmailVerified,
			AuthProvider:  entity.AuthProviderGoogle,
		}
		_, err := u.gateway.CreateUser(user)
		if err != nil {
			return err
		}
		return nil
	}
	if oldUser.AuthProvider.Contains(entity.AuthProviderGoogle) {
		return nil
	}
	oldUser.AuthProvider |= entity.AuthProviderGoogle
	if !oldUser.EmailVerified && options.EmailVerified {
		oldUser.EmailVerified = true
	}
	if oldUser.FamilyName == "" {
		oldUser.FamilyName = options.FamilyName
	}
	if oldUser.GivenName == "" {
		oldUser.GivenName = options.GivenName
	}
	if oldUser.Name == "" {
		oldUser.Name = options.Name
	}
	if oldUser.Picture == "" {
		oldUser.Picture = options.Picture
	}
	_, errUpdate := u.gateway.UpdateUser(*oldUser)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (u *UserUsecase) LoginByProvider(options LoginOptions) error {
	switch options.Provider {
	case entity.AuthProviderLocal:
		return u.Login(options.LocalProvider.Email, options.LocalProvider.Password)
	case entity.AuthProviderGoogle:
		return u.LoginByGoogle(options.GoogleProvider)
	}
	return entity.ErrLoginProviderAuthProviderUknown
}

type GoogleProviderOptions struct {
	Email         string
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	EmailVerified bool
}

type LoginOptions struct {
	Provider      entity.AuthProvider
	LocalProvider struct {
		Email    string
		Password string
	}
	GoogleProvider GoogleProviderOptions
}
