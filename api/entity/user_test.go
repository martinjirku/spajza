package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_UserCreationWithoutError(t *testing.T) {
	_, err := entity.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
	if err != nil {
		t.Errorf("Expected no error")
	}
}
func Test_UserCreationNewEmail(t *testing.T) {
	user, _ := entity.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
	if user.Email != "martinjirku@gmail.com" {
		t.Errorf("Expected no error")
	}
}
func Test_UserCreationWithHashedPassword(t *testing.T) {
	user, _ := entity.NewUserWithPassword("martinjirku@gmail.com", "totalneDobreHeslo")
	if user.Password == "" || user.Password == "totalneDobreHeslo" {
		t.Errorf("Expected password cannot be empty or original value")
	}
}

func Test_PasswordValidation(t *testing.T) {
	password := "totalneDobreHeslo"
	user, _ := entity.NewUserWithPassword("martinjirku@gmail.com", password)
	t.Run("should successfully validate password", func(t *testing.T) {
		if !user.VerifyPassword(password) {
			t.Errorf("Should verify password")
		}
	})
	t.Run("should not successfully validate password", func(t *testing.T) {
		if user.VerifyPassword(password + "asdf") {
			t.Errorf("Should not verify password")
		}
	})
}
