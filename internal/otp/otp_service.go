package otp

import (
	"context"
	"time"

	db "github.com/Dev-Awaab/go_job_portal_api/db/sqlc"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/utils"
)

type otpServive struct {
	repo OtpRepository
	timeout time.Duration
}


func NewOtpService(repo OtpRepository) OtpService {
	return &otpServive{
		repo: repo,
		timeout: time.Duration(5) * time.Second,
	}
}


// Add implements OtpService.
func (o *otpServive) Add(c context.Context, data OtpCreate) (*OtpCode, error) {
	ctx, cancle := context.WithTimeout(c, o.timeout)
	defer cancle()


	ExpiresAt := time.Now().Add(1 * time.Hour)
	

	code := utils.GenerateOtp(6)

	result, err := o.repo.CreateOtp(ctx, db.CreateOtpParams{
		Code:      code,
		Model:     data.Model,
		ModelID:   data.ModelID,
		ExpiresAt: ExpiresAt,
	})

	if err != nil {
		return nil, err
	}

	return &OtpCode{Code: result.Code}, nil
}

// Get implements OtpService.
func (o *otpServive) Get(c context.Context, filter OtpFilter) (*OtpResponse, error) {
	ctx, cancle := context.WithTimeout(c, o.timeout)
	defer cancle()

	result, err := o.repo.GetOtp(ctx, db.GetOtpParams{
		Code:  filter.Code,
		Model: filter.Model,
	})
	if err != nil {
		return nil, err
	}

	return &OtpResponse{
		ID:        result.ID,
		Code:      result.Code,
		Model:     result.Model,
		ModelID:   result.ModelID,
		ExpiresAt: result.ExpiresAt,
		CreatedAt: result.CreatedAt,
	}, nil
}

// Remove implements OtpService.
func (o *otpServive) Remove(c context.Context, filter OtpFilter) error {
	ctx, cancle := context.WithTimeout(c, o.timeout)
	defer cancle()

	result, err := o.repo.GetOtp(ctx, db.GetOtpParams{
		Code:  filter.Code,
		Model: filter.Model,
	})
	if err != nil {
		return err
	}

	return o.repo.DeleteOtpByID(ctx, result.ID)
}
