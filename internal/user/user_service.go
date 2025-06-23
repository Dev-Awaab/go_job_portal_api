package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/utils"
)

type userService struct {
	repo    UserRepository
	timeout time.Duration
}


func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo:    repo,
		timeout: time.Duration(5) * time.Second,
	}
}

// Create implements UserService.
func (u *userService) Register(c context.Context, req *CreateUserReq) (*User, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	_, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err // actual database error
	}

	hashedPassword, err := utils.Hash(req.Password)

	if err != nil {
		return nil, err
	}

	user, err := u.repo.CreateUser(ctx, db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
		Email:     req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// VerifyUserEmail implements UserService.
func (u *userService) UpdateUser(c context.Context, req *UpdateUserReq) (*User, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	params := db.UpdateUserParams{
		ID:              req.ID,
		FirstName:       utils.ToNullString(req.FirstName),
		LastName:        utils.ToNullString(req.LastName),
		Phone:           utils.ToNullString(req.Phone),
		Country:         utils.ToNullString(req.Country),
		City:            utils.ToNullString(req.City),
		AvatarUrl:       utils.ToNullString(req.AvatarUrl),
		IsEmailVerified: utils.ToNullBool(req.IsEmailVerified),
	}

	updatedUser, err := u.repo.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:              updatedUser.ID,
		FirstName:       updatedUser.FirstName,
		LastName:        updatedUser.LastName,
		Email:           updatedUser.Email,
		Password:        updatedUser.Password,
		Phone:           updatedUser.Phone.String,
		Country:         updatedUser.Country.String,
		City:            updatedUser.City.String,
		AvatarUrl:       updatedUser.AvatarUrl.String,
		IsEmailVerified: updatedUser.IsEmailVerified,
		CreatedAt:       updatedUser.CreatedAt,
		UpdatedAt:       updatedUser.UpdatedAt,
		DeletedAt:       updatedUser.DeletedAt.Time,
		Deleted:         updatedUser.Deleted,
	}, nil
}


// Login implements UserService.
func (u *userService) Login(c context.Context, req *LoginUserReq) (*User, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout)
	defer cancel()

	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsEmailVerified {
		return nil, errors.New("confirm your mail")
	}

	isValidPassword := utils.Valid(user.Password, req.Password)
	if !isValidPassword {
		return nil, errors.New("invalid credentials")
	}

	// Do not return the password
	user.Password = ""

	return &User{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Phone:           user.Phone.String,
		Country:         user.Country.String,
		City:            user.City.String,
		AvatarUrl:       user.AvatarUrl.String,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt.Time,
		Deleted:         user.Deleted,
	}, nil
}
