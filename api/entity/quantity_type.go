package entity

import (
	"errors"
	"fmt"
)

type QuantityType string

const (
	QuantityUnknown     QuantityType = "unknown"
	QuantityMass        QuantityType = "mass"        // kg
	QuantityLength      QuantityType = "length"      // m
	QuantityVolume      QuantityType = "volume"      // l
	QuantityTemperature QuantityType = "temperature" // C
	QuantityTime        QuantityType = "time"        // m
	QuantityCount       QuantityType = "count"       // ks
	// AREA        UnitCategory = "area"      // m2
)

var (
	quantities = []QuantityType{QuantityMass, QuantityLength, QuantityVolume, QuantityTemperature, QuantityTime, QuantityCount}
)

func (ct *QuantityType) Scan(value interface{}) error {

	switch typ := value.(type) {
	case *string:
		*ct = QuantityType(value.(string))
	case string:
		*ct = QuantityType(value.(string))
	case []byte:
		*ct = QuantityType(string(value.([]byte)))
	default:
		return fmt.Errorf("could not scan value quantity: %s", typ)
	}
	return ct.IsValid()
}

func (ct *QuantityType) IsValid() error {
	for _, item := range quantities {
		if item == *ct {
			return nil
		}
	}
	return errors.New("unknown quantity")
}

func (ct QuantityType) Value() (string, error) {
	return string(ct), nil
}
