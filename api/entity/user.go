package entity

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthProvider int

const (
	AuthProviderUnknown AuthProvider = iota
	AuthProviderLocal   AuthProvider = 1 << iota
	AuthProviderGoogle
)

func (a AuthProvider) Contains(provider AuthProvider) bool {
	if AuthProviderUnknown == provider {
		return true
	}
	return (a & provider) > 0
}

type User struct {
	ID            int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Password      string
	Email         string
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	EmailVerified bool
	AuthProvider  AuthProvider
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func isPasswordValid(p string) bool {
	return len(p) > 3
}

func NewUser(email string) (User, error) {
	user := User{Email: email}
	if !isEmailValid(email) {
		return user, errors.New("InvalidEmail")
	}
	return user, nil
}

func NewUserWithPassword(email string, password string) (User, error) {
	user := User{Email: email}
	if !isEmailValid(email) {
		return user, errors.New("InvalidEmail")
	}
	if !isPasswordValid(password) {
		return user, errors.New("InvalidPassword")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)
	return user, nil
}

func (u *User) VerifyPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}
