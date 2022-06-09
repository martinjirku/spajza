package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255); index:idx_user_email,unique"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func isPasswordValid(p string) bool {
	return len(p) > 3
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
