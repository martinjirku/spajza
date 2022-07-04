package users

import (
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{db}
}

func (r *UserService) ListAll() ([]*User, error) {
	var u []*User
	err := r.db.Find(&u).Error
	return u, err
}

func (r *UserService) Register(email string, password string) (User, error) {
	user, err := NewUserWithPassword(email, password)
	if err != nil {
		return user, err
	}
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *UserService) Login(email string, password string) error {
	var user User
	result := r.db.Find(&user, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("WrongUsername")
	}
	if !user.VerifyPassword(password) {
		return errors.New("WrongPassword")
	}
	return nil
}
