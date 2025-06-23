package mail

import (
	"fmt"

	"github.com/Dev-Awaab/go_job_portal_api/config"
	"github.com/resend/resend-go/v2"
)


type ResendMailer struct {
	client *resend.Client
	from   string
}

func NewResendMailer(cfg config.Config) (*ResendMailer, error) {
	 if cfg.ResendAPIKey == "" {
		return nil, fmt.Errorf("resend API key is required")
	 }

	 return &ResendMailer{
	client: resend.NewClient(cfg.ResendAPIKey),
	from:   "onboarding@resend.dev",
}, nil
}

func (m *ResendMailer) Send(e Email) error {
	params := &resend.SendEmailRequest{
		To:      []string{e.To},
		From:    m.from,
		Subject: e.Subject,
		Text:    e.TextContent,
		Html:    e.HTMLContent,
	}
	resp, err := m.client.Emails.Send(params)
    if err != nil {
        return fmt.Errorf("resend API error: %w", err)
    }

    if resp.Id == "" {
        return fmt.Errorf("empty response ID from Resend")
    }

    return nil
}