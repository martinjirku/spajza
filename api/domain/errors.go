package domain

import "errors"

var (
	ErrorWrongUsername  = errors.New("wrong username")
	ErrorWrongPassword  = errors.New("wrong password")
	ErrorNothingUpdated = errors.New("nothing update")
	ErrorWrongParameter = errors.New("invalid parameter")
	ErrorWrongField     = errors.New("invalid field")
)
