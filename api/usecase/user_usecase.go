package usecase

import (
	"github.com/martinjirku/zasobar/entity"
)

type UserUsecase struct {
	gateway Gateway
}

type Gateway interface {
	Login(email string, password string) error
	Register(email string, password string) (entity.User, error)
}

func GetUserUsecase(registrator Gateway) UserUsecase {
	return UserUsecase{registrator}
}

func (u *UserUsecase) Register(email, password string) (entity.User, error) {
	return u.gateway.Register(email, password)
}

func (u *UserUsecase) Login(email, password string) error {
	return u.gateway.Login(email, password)
}
