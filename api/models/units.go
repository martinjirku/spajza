package models

import (
	"gorm.io/gorm"
)

type UnitCategory string

const (
	UNKNOWN     UnitCategory = "unknown"
	MASS        UnitCategory = "mass"        // kg
	LENGTH      UnitCategory = "length"      // m
	VOLUME      UnitCategory = "volume"      // l
	TEMPERATURE UnitCategory = "temperature" // C
	TIME        UnitCategory = "time"        // m
	COUNT       UnitCategory = "count"       // ks
	// AREA        UnitCategory = "area"        // m2
)

func (ct *UnitCategory) Scan(value interface{}) error {
	*ct = UnitCategory(value.([]byte))
	return nil
}

func (ct UnitCategory) Value() (string, error) {
	return string(ct), nil
}

type Unit struct {
	gorm.Model
	Title    string       `gorm:"type:varchar(255)"`
	UnitName string       `gorm:"type:varchar(50)"`
	Category UnitCategory `gorm:"type:varchar(255)"`
}
