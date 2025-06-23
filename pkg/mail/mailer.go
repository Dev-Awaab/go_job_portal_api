package mail

import (
	"fmt"

	"github.com/Dev-Awaab/go_job_portal_api/config"
)

type Mailer interface {
	Send(email Email) error
}

// NewMailer creates a mailer for the specified provider
func NewMailer(provider Provider, cfg config.Config) (Mailer, error) {
	switch provider {
	case Resend:
		return NewResendMailer(cfg)
	default:
		return nil, fmt.Errorf("unsupported mail provider: %d", provider)
	}
}