// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
    username,
    hash_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
)
RETURNING username, hash_password, full_name, email, password_change_at, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hash_password, full_name, email, password_change_at, created_at FROM "user"
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreatedAt,
	)
	return i, err
}