package mail

type Provider int

const (
	Resend Provider = iota
	SendGrid
	SES
	SMTP
)

type Email struct {
	To          string
	Subject     string
	HTMLContent string
	TextContent string
}
// package mail

// type Provider string

// const (
// 	ProviderSendGrid Provider = "sendgrid"
// 	ProviderSES      Provider = "ses"
// 	ProviderSMTP     Provider = "smtp"
// 	ProviderResend   Provider = "resend"
// )

// func (p Provider) String() string {
// 	return string(p)
// }

// func (p Provider) IsValid() bool {
// 	switch p {
// 	case ProviderSendGrid, ProviderSES, ProviderSMTP, ProviderResend:
// 		return true
// 	default:
// 		return false
// 	}
// }
// type Email struct {
// 	To 		string
// 	Subject string
// 	HTMLContent string
// 	TextContent string
// }


// type MailConfig struct {
//     Provider   string
//     FromEmail  string
//     FromName   string
//     APIKey     string 
//     Region     string 
//     SMTPHost   string 
//     SMTPPort   int    
//     SMTPUser   string 
//     SMTPPass   string 
// }