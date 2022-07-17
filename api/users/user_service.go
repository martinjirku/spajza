package users

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{db}
}

func (r *UserService) ListAll() ([]*User, error) {
	var users []*User
	selectAllUsersStmt := "SELECT id, created_at, updated_at, deleted_at, password, email FROM users"
	rows, err := r.db.Query(selectAllUsersStmt)
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

func (r *UserService) Register(ctx context.Context, email string, password string) (User, error) {
	user, err := NewUserWithPassword(email, password)
	if err != nil {
		return user, err
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res, err := r.db.ExecContext(ctx, "INSERT INTO users(created_at, updated_at, email, password) VALUES (?,?,?,?)",
		user.CreatedAt, user.UpdatedAt, user.Email, user.Password)
	if err != nil {
		return user, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}
	user.ID = uint(id)
	return user, nil
}

func (r *UserService) Login(ctx context.Context, email string, password string) error {
	var user = User{}
	err := r.db.
		QueryRowContext(ctx, "SELECT id, created_at, updated_at, password, email FROM users WHERE email=? AND deleted_at IS NULL", email).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Password, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		return ErrorWrongUsername()
	case err != nil:
		return err
	case !user.VerifyPassword(password):
		return ErrorWrongPassword()
	default:
		return nil
	}
}
