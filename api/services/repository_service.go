package services

import (
	"github.com/martinjirku/zasobar/users"
	"gorm.io/gorm"
)

type RepositoryService struct {
	User users.UserService
}

func NewRepositoryService(db *gorm.DB) RepositoryService {
	return RepositoryService{User: users.NewUserService(db)}
}
