package otp

import (
	"context"
	"time"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
)


type Otp struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Model     string    `json:"model"`
	ModelID   string    `json:"model_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type OtpRepository interface {
	CreateOtp(ctx context.Context, params db.CreateOtpParams) (*db.Otp, error)
	GetOtpByID(ctx context.Context, id int64) (*db.Otp, error)
	GetOtpByCodeAndModel(ctx context.Context, code string, model string, modelID string) (*db.Otp, error)
	DeleteOtpByID(ctx context.Context, id int64) error
	DeleteExpiredOtps(ctx context.Context) error
	GetOtp(ctx context.Context, params db.GetOtpParams) (*db.Otp, error)
}

type OtpService interface {
	Add(ctx context.Context, data OtpCreate) (*OtpCode, error)
	Get(ctx context.Context, filter OtpFilter) (*OtpResponse, error)
	Remove(ctx context.Context, filter OtpFilter) error

}

type OtpCreate struct {
	Model    string
	ModelID  string
}

type OtpCode struct {
	Code string
	
}


type OtpFilter struct {
	Code    string
	Model   string
	// ModelID *string
}

type OtpResponse struct {
	ID        int64
	Code      string
	Model     string
	ModelID   string
	ExpiresAt time.Time
	CreatedAt time.Time
}
