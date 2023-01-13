package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/martinjirku/zasobar/entity"
)

type UserRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserRepository(ctx context.Context, db *sql.DB) *UserRepository {
	return &UserRepository{db, ctx}
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	userStmt := "SELECT id, created_at, updated_at, deleted_at, password, email FROM users where email = ?"
	row := r.db.QueryRowContext(r.ctx, userStmt, email)
	if row.Err() != nil {
		return nil, row.Err()
	}
	user := entity.User{}
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListAll() ([]*entity.User, error) {
	var users []*entity.User
	selectAllUsersStmt := "SELECT id, created_at, updated_at, deleted_at, password, email FROM users"
	rows, err := r.db.QueryContext(r.ctx, selectAllUsersStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = entity.User{}
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Password, &user.Email); err != nil {
			fmt.Printf("could not scan row: %v", err)
		}

		users = append(users, &user)
	}
	return users, err
}

func (r *UserRepository) CreateUser(user entity.User) (*entity.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res, err := r.db.ExecContext(r.ctx, "INSERT INTO users(created_at,updated_at,email,password,name,given_name,family_name,picture,email_verified,auth_provider) VALUES (?,?,?,?,?,?,?,?,?,?)",
		user.CreatedAt, user.UpdatedAt, user.Email, user.Password, user.Name, user.GivenName, user.FamilyName, user.Picture, user.EmailVerified, user.AuthProvider)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = uint(id)
	return &user, nil
}

func (r *UserRepository) UpdateUser(user entity.User) (*entity.User, error) {
	user.UpdatedAt = time.Now()
	stmt := "UPDATE users SET updated_at=?,email=?,password=?,name=?,family_name=?,given_name=?,email_verified=?,auth_provider=?,picture=? WHERE ID=?"
	_, err := r.db.ExecContext(r.ctx, stmt, user.UpdatedAt, user.Email, user.Password, user.Name,
		user.FamilyName, user.GivenName, user.EmailVerified, user.AuthProvider, user.Picture, user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Register(email string, password string) (*entity.User, error) {
	user, err := entity.NewUserWithPassword(email, password)
	if err != nil {
		return nil, err
	}
	return r.CreateUser(user)
}

func (r *UserRepository) Login(email string, password string) error {
	var user = entity.User{}
	err := r.db.
		QueryRowContext(r.ctx, "SELECT id, created_at, updated_at, password, email FROM users WHERE email=? AND deleted_at IS NULL", email).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Password, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		return entity.ErrWrongUsername
	case err != nil:
		return err
	case !user.VerifyPassword(password):
		return entity.ErrWrongPassword
	default:
		return nil
	}
}
