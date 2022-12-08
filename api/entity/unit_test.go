package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_UnitValidation(t *testing.T) {
	t.Run("valid unit and quantity", func(t *testing.T) {
		unit := entity.Unit{Name: entity.UnitGram, Quantity: entity.QuantityMass}
		isValid := unit.Validate()
		if !isValid {
			t.Error("Expected valid, but received invalid")
		}
	})
	t.Run("invalid unit", func(t *testing.T) {
		unit := entity.Unit{Name: "test"}
		isValid := unit.Validate()
		if isValid {
			t.Error("Expected invalid, but received valid")
		}
	})
}
