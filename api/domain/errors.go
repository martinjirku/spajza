package domain

import "errors"

var (
	ErrorWrongUsername = errors.New("wrong username")
	ErrorWrongPassword = errors.New("wrong password")
)
