package domain_test

import (
	"testing"

	"github.com/martinjirku/zasobar/domain"
)

func TestQuantityScanString(t *testing.T) {
	massStr := "mass"
	var quantity domain.Quantity
	err := quantity.Scan(massStr)
	if err != nil {
		t.Errorf("During scan error occured `%s`", err.Error())
	}
	if quantity != domain.MASS {
		t.Errorf("Expected `%s`, but received `%s`", domain.MASS, quantity)
	}
}

func TestQuantityScanInvalidString(t *testing.T) {
	massStr := "masses"
	var quantity domain.Quantity
	err := quantity.Scan(massStr)
	if err == nil {
		t.Errorf("Expected error, but no error returned")
	}
}
