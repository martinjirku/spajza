package entity

import "errors"

var (
	ErrWrongUsername  = errors.New("wrongUsername")
	ErrWrongPassword  = errors.New("wrongPassword")
	ErrNothingUpdated = errors.New("nothingUpdate")
	ErrWrongParameter = errors.New("invalidParameter")
	ErrWrongField     = errors.New("invalidField")
)
