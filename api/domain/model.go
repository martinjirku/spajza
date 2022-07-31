package domain

import (
	"errors"
	"fmt"
)

type Quantity string

const (
	MASS        Quantity = "mass"        // kg
	LENGTH      Quantity = "length"      // m
	VOLUME      Quantity = "volume"      // l
	TEMPERATURE Quantity = "temperature" // C
	TIME        Quantity = "time"        // m
	COUNT       Quantity = "count"       // ks
	// AREA        UnitCategory = "area"      // m2
)

var (
	quantities = []Quantity{MASS, LENGTH, VOLUME, TEMPERATURE, TIME, COUNT}
)

func (ct *Quantity) Scan(value interface{}) error {
	switch t := value.(type) {
	case *string:
		*ct = Quantity(value.(string))
	case []byte:
		*ct = Quantity(string(value.([]byte)))
	default:
		return fmt.Errorf("could not scan value quantity: %s", t)
	}
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
