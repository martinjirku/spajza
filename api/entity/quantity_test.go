package entity_test

import (
	"testing"

	"github.com/martinjirku/zasobar/entity"
)

func Test_QuantityScanString(t *testing.T) {
	massStr := "mass"
	var quantity entity.Quantity
	err := quantity.Scan(massStr)
	if err != nil {
		t.Errorf("During scan error occured `%s`", err.Error())
	}
	if quantity != entity.MASS {
		t.Errorf("Expected `%s`, but received `%s`", entity.MASS, quantity)
	}
}

func Test_QuantityScanInvalidString(t *testing.T) {
	massStr := "masses"
	var quantity entity.Quantity
	err := quantity.Scan(massStr)
	if err == nil {
		t.Errorf("Expected error, but no error returned")
	}
}
