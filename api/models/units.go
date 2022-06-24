package models

import (
	"gorm.io/gorm"
)

type Quantity string

const (
	UNKNOWN     Quantity = "unknown"
	MASS        Quantity = "mass"        // kg
	LENGTH      Quantity = "length"      // m
	VOLUME      Quantity = "volume"      // l
	TEMPERATURE Quantity = "temperature" // C
	TIME        Quantity = "time"        // m
	COUNT       Quantity = "count"       // ks
	// AREA        UnitCategory = "area"        // m2
)

func (ct *Quantity) Scan(value interface{}) error {
	*ct = Quantity(value.([]byte))
	return nil
}

func (ct Quantity) Value() (string, error) {
	return string(ct), nil
}

type Unit struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(250)"`
	Quantity Quantity `gorm:"type:varchar(25)"`
	Name     string   `gorm:"type:varchar(50)"`
	Symbol   string   `gorm:"type:varchar(10)"`
}
