package services

import (
	"database/sql"

	"github.com/martinjirku/zasobar/users"
	"gorm.io/gorm"
)

type RepositoryService struct {
	User users.UserService
}

func NewRepositoryService(db *gorm.DB, rawDb *sql.DB) RepositoryService {
	return RepositoryService{User: users.NewUserService(db, rawDb)}
}
