package models

import "errors"

type Quantity string

const (
	MASS        Quantity = "mass"        // kg
	LENGTH      Quantity = "length"      // m
	VOLUME      Quantity = "volume"      // l
	TEMPERATURE Quantity = "temperature" // C
	TIME        Quantity = "time"        // m
	COUNT       Quantity = "count"       // ks
	// AREA        UnitCategory = "area"        // m2
)

var (
	quantities = []Quantity{MASS, LENGTH, VOLUME, TEMPERATURE, TIME, COUNT}
)

func (ct *Quantity) Scan(value interface{}) error {
	*ct = Quantity(value.(string))
	return ct.IsValid()
}

func (ct *Quantity) IsValid() error {
	for _, item := range quantities {
		if item == *ct {
			return nil
		}
	}
	return errors.New("unknown quantity")
}

func (ct Quantity) Value() (string, error) {
	return string(ct), nil
}
