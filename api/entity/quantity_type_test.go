package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_QuantityTypeScanString(t *testing.T) {
	massStr := "mass"
	var quantity entity.QuantityType
	err := quantity.Scan(massStr)
	if err != nil {
		t.Errorf("During scan error occured `%s`", err.Error())
	}
	if quantity != entity.QuantityMass {
		t.Errorf("Expected `%s`, but received `%s`", entity.QuantityMass, quantity)
	}
}

func Test_QuantityTypeScanInvalidString(t *testing.T) {
	massStr := "masses"
	var quantity entity.QuantityType
	err := quantity.Scan(massStr)
	if err == nil {
		t.Errorf("Expected error, but no error returned")
	}
}
