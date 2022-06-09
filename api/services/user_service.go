package services

import (
	"errors"

	models "github.com/martinjirku/zasobar/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (r *UserService) ListAll() ([]*models.User, error) {
	var u []*models.User
	err := r.db.Find(&u).Error
	return u, err
}

func (r *UserService) Register(email string, password string) (*models.User, error) {
	user, err := models.NewUserWithPassword(email, password)
	if err != nil {
		return nil, err
	}
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *UserService) Login(email string, password string) error {
	var user = models.User{Email: email}
	result := r.db.First(&user)
	if result.Error != nil {
		return result.Error
	}
	if !user.VerifyPassword(password) {
		return errors.New("WrongPassword")
	}
	return nil
}
