package services

import "gorm.io/gorm"

type RepositoryService struct {
	User UserService
}

func NewRepositoryService(db *gorm.DB) RepositoryService {
	return RepositoryService{User: UserService{db}}
}
