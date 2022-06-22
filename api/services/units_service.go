package services

import (
	u "github.com/bcicen/go-units"
)

func PossibleUnits() []u.Unit {
	u.NewUnit("cup", "cup", u.Volume)
	return u.All()
}
