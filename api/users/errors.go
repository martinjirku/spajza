package users

import "errors"

const (
	wrongUsername = "WrongUsername"
	wrongPassword = "WrongPassword"
)

func ErrorWrongUsername() error {
	return errors.New(wrongUsername)
}

func ErrorWrongPassword() error {
	return errors.New(wrongPassword)
}
