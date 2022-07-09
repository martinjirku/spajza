package categories

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title       string `gorm:"type:varchar(250)"`
	Path        string `gorm:"type:varchar(250)"`
	DefaultUnit string `gorm:"type:varchar(50)"`
}
