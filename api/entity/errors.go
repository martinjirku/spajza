package entity

import "errors"

var (
	ErrWrongUsername                   = errors.New("wrongUsername")
	ErrWrongPassword                   = errors.New("wrongPassword")
	ErrNothingUpdated                  = errors.New("nothingUpdate")
	ErrInvalidEntity                   = errors.New("entityInvalid")
	ErrEntityNotFound                  = errors.New("entityNotFound")
	ErrEntityConflict                  = errors.New("entityConflict")
	ErrInvalidParameter                = errors.New("invalidParameter")
	ErrInvalidField                    = errors.New("invalidField")
	ErrLoginProviderAuthProviderUknown = errors.New("user.AuthProviderUknown")
	ErrStorageItemNotEnoughQuantity    = errors.New("storageItem.NotEnoughQuantity")
)
