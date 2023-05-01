// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package generated

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (phone, first_name, last_name)
VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Phone, arg.FirstName, arg.LastName)
	return err
}

const createUserCode = `-- name: CreateUserCode :one
INSERT INTO users_codes(user_id, code)
VALUES ($1, $2)
RETURNING id
`

type CreateUserCodeParams struct {
	UserID int32 `json:"user_id"`
	Code   int32 `json:"code"`
}

func (q *Queries) CreateUserCode(ctx context.Context, arg CreateUserCodeParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUserCode, arg.UserID, arg.Code)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, phone, first_name, last_name, created_at
FROM users
WHERE phone = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserCode = `-- name: GetUserCode :one
SELECT id, user_id, code, created_at
FROM users_codes
WHERE user_id = $1
ORDER BY created_at DESC LIMIT 1
`

func (q *Queries) GetUserCode(ctx context.Context, userID int32) (UsersCode, error) {
	row := q.db.QueryRowContext(ctx, getUserCode, userID)
	var i UsersCode
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Code,
		&i.CreatedAt,
	)
	return i, err
}
