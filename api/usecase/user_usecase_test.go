package usecase_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
	"github.com/martinjirku/zasobar/mocks"
	"github.com/martinjirku/zasobar/usecase"
	"github.com/stretchr/testify/mock"
)

func Test_UserUsecase_LoginByGoogle(t *testing.T) {
	t.Run("should register when user does not exists", func(t *testing.T) {
		userGateway := mocks.NewUserGateway(t)
		userGateway.EXPECT().GetByEmail("test@gmail.com").Return(nil, entity.ErrEntityNotFound)
		userGateway.EXPECT().CreateUser(mock.MatchedBy(func(u entity.User) bool {
			if u.Password != "" {
				return false
			}
			if u.Email != "test@gmail.com" {
				return false
			}
			if u.Name != "John Doe" {
				return false
			}
			if u.GivenName != "John" {
				return false
			}
			if u.FamilyName != "Doe" {
				return false
			}
			if !u.AuthProvider.Contains(entity.AuthProviderGoogle) {
				return false
			}
			return true
		})).Return(&entity.User{
			ID:            1,
			Email:         "test@gmail.com",
			Name:          "John Doe",
			GivenName:     "John",
			FamilyName:    "Doe",
			Picture:       "http://url.to/image",
			EmailVerified: false,
			AuthProvider:  entity.AuthProviderGoogle,
		}, nil)
		options := usecase.GoogleProviderOptions{
			Email:         "Test@gmail.com",
			Name:          "John Doe",
			GivenName:     "John",
			FamilyName:    "Doe",
			Picture:       "http://url.to/image",
			EmailVerified: true,
		}
		userUsecase := usecase.GetUserUsecase(userGateway)
		err := userUsecase.LoginByGoogle(options)
		if err != nil {
			t.Errorf("should not fail, but received %s", err)
		}
	})
	t.Run("should update only empty fields, set verified when verified, and set authProvider when only AuthLocal exists", func(t *testing.T) {
		userGateway := mocks.NewUserGateway(t)
		oldUser := &entity.User{
			ID:            1,
			Email:         "test@gmail.com",
			Picture:       "http://url.to.old/image",
			EmailVerified: false,
			AuthProvider:  entity.AuthProviderLocal,
		}
		userGateway.EXPECT().GetByEmail("test@gmail.com").Return(oldUser, nil)
		userGateway.EXPECT().UpdateUser(mock.MatchedBy(func(u entity.User) bool {
			if u.ID != 1 ||
				u.Password != "" ||
				u.Email != "test@gmail.com" ||
				u.Name != "John Doe" ||
				u.GivenName != "John" ||
				u.FamilyName != "Doe" ||
				!u.AuthProvider.Contains(entity.AuthProviderGoogle) ||
				!u.AuthProvider.Contains(entity.AuthProviderLocal) ||
				u.Picture != "http://url.to.old/image" ||
				!u.EmailVerified {
				return false
			}
			return true
		})).Return(&entity.User{
			ID:            1,
			Email:         "test@gmail.com",
			Name:          "John Doe",
			GivenName:     "John",
			FamilyName:    "Doe",
			Picture:       "http://url.to.old/image",
			EmailVerified: true,
			AuthProvider:  entity.AuthProviderGoogle,
		}, nil)
		options := usecase.GoogleProviderOptions{
			Email:         "Test@gmail.com",
			Name:          "John Doe",
			GivenName:     "John",
			FamilyName:    "Doe",
			Picture:       "http://url.to.new/image",
			EmailVerified: true,
		}
		userUsecase := usecase.GetUserUsecase(userGateway)
		err := userUsecase.LoginByGoogle(options)
		if err != nil {
			t.Errorf("should not fail, but received %s", err)
		}
	})
}
