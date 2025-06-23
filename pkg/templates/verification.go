package templates

import "html/template"

var OTPEmailTemplate = template.Must(template.New("otpEmail").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>
        .container { max-width: 600px; margin: 0 auto; font-family: Arial, sans-serif; }
        .header { background-color: #2563eb; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; }
        .otp-code { 
            font-size: 24px; 
            font-weight: bold; 
            letter-spacing: 2px; 
            text-align: center;
            margin: 20px 0;
            color: #2563eb;
        }
        .footer { margin-top: 20px; font-size: 12px; color: #6b7280; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Your OTP Code</h1>
        </div>
        <div class="content">
            <p>Hello {{.Name}},</p>
            <p>Here is your one-time password (OTP) for verification:</p>
            <div class="otp-code">{{.OTP}}</div>
            <p>This code will expire in {{.ExpiryMinutes}} minutes.</p>
            <p>If you didn't request this, please ignore this email.</p>
        </div>
        <div class="footer">
            <p>Best regards,<br>The {{.AppName}} Team</p>
        </div>
    </div>
</body>
</html>
`))

type OTPEmailData struct {
    Name          string
    OTP           string
    ExpiryMinutes int
    AppName       string
}