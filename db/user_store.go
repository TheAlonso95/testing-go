package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goschool/crud/types"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	GetUser(ctx context.Context, email string) (*types.User, error)
}

type SQLiteUserStore struct {
	db *sql.DB
}

func NewSQLiteUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{
		db: db,
	}
}

func (u *SQLiteUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	query := `INSERT INTO users (name, email, password_hash)
	VALUES (?, ?, ?)
	RETURNING id`

	var userId string
	err := u.db.QueryRowContext(ctx, query, user.Name, user.Email, user.PasswordHash).Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("Create User: %w", err)
	}

	// Not sure if I agree with this approach
	// but for the sake of the course let's keep it like this
	user.ID = userId

	return user, nil
}

func (u *SQLiteUserStore) GetUser(ctx context.Context, email string) (*types.User, error) {
	var user types.User

	query := `SELECT * FROM users WHERE email = ?`

	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.Name, &user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch User: %w", err)
	}

	return &user, nil
}
