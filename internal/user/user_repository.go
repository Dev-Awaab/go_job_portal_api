package user

import (
	"context"
	"database/sql"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
)

type userRepository struct {
	queries *db.Queries
}



func NewUserRepository(dbConn *sql.DB) UserRepository {
	return &userRepository{
		queries: db.New(dbConn),
	}
}

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (*db.User, error) {
	user, err := u.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := u.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID implements UserRepository.
func (u *userRepository) GetUserByID(ctx context.Context, id int64) (*db.User, error) {
	user, err := u.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


// UpdateUser implements UserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, params db.UpdateUserParams) (*db.User, error) {
	user, err := u.queries.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}