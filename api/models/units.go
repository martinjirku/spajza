package models

import (
	"gorm.io/gorm"
)

type UnitCategory string

const (
	MASS   UnitCategory = "MASS"   // kg
	LENGTH UnitCategory = "LENGTH" // m
	AREA   UnitCategory = "AREA"   // m2
	VOLUME UnitCategory = "VOLUME" // m3
	COUNT  UnitCategory = "COUNT"  // ks
	TIME   UnitCategory = "TIME"   // m
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
	UnitName string       `gorm:"type:varchar(30)"`
	Category UnitCategory `gorm:"type:varchar(255)"`
}
