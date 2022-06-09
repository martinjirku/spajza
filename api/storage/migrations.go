package storage

import (
	"github.com/martinjirku/zasobar/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
