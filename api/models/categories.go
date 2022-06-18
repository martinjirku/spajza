package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `gorm:"type:varchar(255)"`
	Unit  Unit
	Path  string `gorm:"type:varchar(255)"`
}
