package services

import (
	"database/sql"

	"github.com/martinjirku/zasobar/users"
)

type RepositoryService struct {
	User users.UserService
}

func NewRepositoryService(db *sql.DB) RepositoryService {
	return RepositoryService{User: users.NewUserService(db)}
}
