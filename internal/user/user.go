package user

import (
	"context"
	"time"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
)

type User struct {
	ID              int64          `json:"id"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	Phone           string		   `json:"phone"`
	Country         string		   `json:"country"`
	City            string		   `json:"city"`
	AvatarUrl       string		   `json:"avatar_url"`
	IsEmailVerified bool           `json:"is_email_verified"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       time.Time  	   `json:"deleted_at"`
	Deleted         bool           `json:"deleted"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, params db.CreateUserParams) (*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	GetUserByID(ctx context.Context, id int64) (*db.User, error)
	UpdateUser(ctx context.Context, params db.UpdateUserParams) (*db.User, error) 
}

type UserService interface {
	Register(c context.Context, req *CreateUserReq)(*User, error)
	Login(c context.Context, req *LoginUserReq)(*User, error)
	UpdateUser(c context.Context, req *UpdateUserReq)(*User, error)

}

type VerifyUserEmailReq struct {
	Code string `json:"code" binding:"required"`
}

type CreateUserReq struct {
	FirstName string `json:"firstName" binding:"required"`
    LastName  string `json:"lastName" binding:"required"`
    Email    string `json:"email" binding:"required,email"` 
    Password string `json:"password" binding:"required,min=6"`
}




type LoginUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}


type UpdateUserReq struct {
	ID              int64   `json:"id" binding:"required"`
	FirstName       *string `json:"first_name,omitempty"`
	LastName        *string `json:"last_name,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	Country         *string `json:"country,omitempty"`
	City            *string `json:"city,omitempty"`
	AvatarUrl       *string `json:"avatar_url,omitempty"`
	IsEmailVerified *bool   `json:"is_email_verified,omitempty"`
}
