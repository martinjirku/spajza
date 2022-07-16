package users

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	db    *gorm.DB
	rawDb *sql.DB
}

func NewUserService(db *gorm.DB, rawDb *sql.DB) UserService {
	return UserService{db, rawDb}
}

func (r *UserService) ListAll() ([]*User, error) {
	var users []*User
	selectAllUsersStmt := "SELECT id, created_at, updated_at, deleted_at, password, email FROM users"
	rows, err := r.rawDb.Query(selectAllUsersStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user = User{}
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Password, &user.Email); err != nil {
			fmt.Printf("could not scan row: %v", err)
		}

		users = append(users, &user)
	}
	return users, err
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
		return ErrorWrongUsername()
	}
	if !user.VerifyPassword(password) {
		return ErrorWrongPassword()
	}
	return nil
}
