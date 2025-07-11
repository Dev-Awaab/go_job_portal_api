// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name,last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id, first_name, last_name, email, password, phone, country, city, avatar_url, is_email_verified, created_at, updated_at, deleted_at, deleted
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Country,
		&i.City,
		&i.AvatarUrl,
		&i.IsEmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Deleted,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, phone, country, city, avatar_url, is_email_verified, created_at, updated_at, deleted_at, deleted FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Country,
		&i.City,
		&i.AvatarUrl,
		&i.IsEmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Deleted,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, email, password, phone, country, city, avatar_url, is_email_verified, created_at, updated_at, deleted_at, deleted FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Country,
		&i.City,
		&i.AvatarUrl,
		&i.IsEmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Deleted,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  first_name = COALESCE($1, first_name),
  last_name = COALESCE($2, last_name),
  phone = COALESCE($3, phone),
  country = COALESCE($4, country),
  city = COALESCE($5, city),
  avatar_url = COALESCE($6, avatar_url),
  is_email_verified = COALESCE($7, is_email_verified),
  updated_at = NOW()
WHERE id = $8
RETURNING id, first_name, last_name, email, password, phone, country, city, avatar_url, is_email_verified, created_at, updated_at, deleted_at, deleted
`

type UpdateUserParams struct {
	FirstName       sql.NullString `json:"first_name"`
	LastName        sql.NullString `json:"last_name"`
	Phone           sql.NullString `json:"phone"`
	Country         sql.NullString `json:"country"`
	City            sql.NullString `json:"city"`
	AvatarUrl       sql.NullString `json:"avatar_url"`
	IsEmailVerified sql.NullBool   `json:"is_email_verified"`
	ID              int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.Country,
		arg.City,
		arg.AvatarUrl,
		arg.IsEmailVerified,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.Country,
		&i.City,
		&i.AvatarUrl,
		&i.IsEmailVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Deleted,
	)
	return i, err
}
