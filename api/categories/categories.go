package categories

import (
	"github.com/martinjirku/zasobar/units"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title        string         `gorm:"type:varchar(250)"`
	Path         string         `gorm:"type:varchar(250)"`
	DefaultUnit  string         `gorm:"type:varchar(50)"`
	QuantityType units.Quantity `gorm:"type:varchar(50)"`
}
