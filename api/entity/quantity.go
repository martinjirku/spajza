package entity

import (
	"errors"
	"fmt"
)

type Quantity string

const (
	QuantityMass        Quantity = "mass"        // kg
	QuantityLength      Quantity = "length"      // m
	QuantityVolume      Quantity = "volume"      // l
	QuantityTemperature Quantity = "temperature" // C
	QuantityTime        Quantity = "time"        // m
	QuantityCount       Quantity = "count"       // ks
	// AREA        UnitCategory = "area"      // m2
)

var (
	quantities = []Quantity{QuantityMass, QuantityLength, QuantityVolume, QuantityTemperature, QuantityTime, QuantityCount}
)

func (ct *Quantity) Scan(value interface{}) error {

	switch typ := value.(type) {
	case *string:
		*ct = Quantity(value.(string))
	case string:
		*ct = Quantity(value.(string))
	case []byte:
		*ct = Quantity(string(value.([]byte)))
	default:
		return fmt.Errorf("could not scan value quantity: %s", typ)
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
