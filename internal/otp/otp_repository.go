package otp

import (
	"context"
	"database/sql"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
)

type otpRepository struct {
	queries *db.Queries
}


func NewOtpRepository(dbConn *sql.DB) OtpRepository {
	return &otpRepository{
		queries: db.New(dbConn),
	}
}

// CreateOtp implements OtpRepository.
func (o *otpRepository) CreateOtp(ctx context.Context, params db.CreateOtpParams) (*db.Otp, error) {
	otp, err := o.queries.CreateOtp(ctx, params)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

// DeleteExpiredOtps implements OtpRepository.
func (o *otpRepository) DeleteExpiredOtps(ctx context.Context) error {
	err := o.queries.DeleteExpiredOtps(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DeleteOtpByID implements OtpRepository.
func (o *otpRepository) DeleteOtpByID(ctx context.Context, id int64) error {
	err := o.queries.DeleteOtpByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetOtpByCodeAndModel implements OtpRepository.
func (o *otpRepository) GetOtpByCodeAndModel(ctx context.Context, code string, model string, modelID string) (*db.Otp, error) {
	otp, err := o.queries.GetOtpByCodeAndModel(ctx, db.GetOtpByCodeAndModelParams{
		Code:    code,
		Model:   model,
		ModelID: modelID,
	})

	if err != nil {
		return nil, err
	}

	return &otp, nil
}

// GetOtpByID implements OtpRepository.
func (o *otpRepository) GetOtpByID(ctx context.Context, id int64) (*db.Otp, error) {
	otp, err := o.queries.GetOtpByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &otp, err
}

// GetOtp implements OtpRepository.
func (o *otpRepository) GetOtp(ctx context.Context, params db.GetOtpParams) (*db.Otp, error) {
	otp, err := o.queries.GetOtp(ctx, params)

	if err != nil {
		return nil, err
	}

	return &otp, err

}
