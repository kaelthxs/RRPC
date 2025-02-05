package repository

import (
	"RRPC"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(users RRPC.Users) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (username, email, password_hash, role, created_at) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(q, users.Username, users.Email, users.Password_hash, users.Role, users.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (RRPC.Users, error) {
	var user RRPC.Users
	q := fmt.Sprintf("SELECT username, password_hash FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, q, username, password)

	return user, err
}
