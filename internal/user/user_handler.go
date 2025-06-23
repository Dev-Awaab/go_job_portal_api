package user

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dev-Awaab/go_job_portal_api/internal/otp"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/mail"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/templates"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/utils"
	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	 UserService
	 mailer mail.Mailer
	 otp otp.OtpService
}

func NewUserHandler(us UserService, mailer mail.Mailer, otp otp.OtpService) *UserHandler {
	return &UserHandler{
		UserService: us,
		mailer:      mailer,
		otp: otp,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req CreateUserReq


	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON or empty body", err, nil)
		return
	}

	if err := utils.ValidateRequest(req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err, nil)
		return
	}



	res, err := h.UserService.Register(c.Request.Context(), &req)
	
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err, nil)
		return
	}
	otpReq := otp.OtpCreate{
		Model: string(utils.User),
		ModelID:  fmt.Sprintf("%v", res.ID),
	}
	otpCode, err := h.otp.Add(c.Request.Context(), otpReq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate OTP", err, nil)
		return
	}

	err = h.sendOTP(res.Email, res.FirstName+" "+res.LastName, otpCode.Code)
	if err != nil {
		// utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to send otp", err, nil)
		fmt.Println("OTP email error:", err)
	}

	utils.SuccessResponse(c, http.StatusCreated, "User Created Check your mail", nil)
}

func (h *UserHandler) VerifyUserEmail(c *gin.Context) {
   var req VerifyUserEmailReq


   if err := c.ShouldBindJSON(&req); err != nil {
	utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON or empty body", err, nil)
	return
}

if err := utils.ValidateRequest(req); err != nil {
	utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err, nil)
	return
}



otpFilter := otp.OtpFilter{
	Model:   string(utils.User),
	Code:    req.Code,  
}


otp, err := h.otp.Get(c.Request.Context(), otpFilter)
if err != nil {
	utils.ErrorResponse(c, http.StatusBadRequest,"Invalid otp", nil, nil)
	return
}


modelID, err := strconv.ParseInt(otp.ModelID, 10, 64)
if err != nil {
	utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid user ID in OTP", err, nil)
	return
}
verified := true
_, err = h.UserService.UpdateUser(c.Request.Context(), &UpdateUserReq{
	ID: modelID,
	IsEmailVerified: &verified,
})
if err != nil {
	utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user verification status", err, nil)
	return
}

utils.SuccessResponse(c, http.StatusOK, "Email verified successfully", nil)
}

func (h *UserHandler) Login(c *gin.Context){
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON or empty body", err, nil)
		return
	}

	if err := utils.ValidateRequest(req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), err, nil)
		return
	}
	


	user, err := h.UserService.Login(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error(), err, nil)
		return
	}

	token, err := utils.GenerateJWT(user.ID, utils.User)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token", err, nil)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	})

}

 



  /*
     * ---------------------------------------------------------------------------------------------------------
     * helper functions
     * ---------------------------------------------------------------------------------------------------------
     */
	func (h *UserHandler) sendOTP(email, name, otpCode string) error {
		// Prepare template data
		data := templates.OTPEmailData{
			Name:          name,
			OTP:           otpCode,
			ExpiryMinutes: int(1),
			AppName:       "Job Portal",
		}
	
		// Render HTML content
		var htmlContent bytes.Buffer
		if err := templates.OTPEmailTemplate.Execute(&htmlContent, data); err != nil {
			return err
		}
	
		// Create email
		emailMsg := mail.Email{
			To:          email,
			Subject:     "Your Verification Code",
			HTMLContent: htmlContent.String(),
			TextContent: h.generatePlainTextOTP(data),
		}
	
		// Send email
		return h.mailer.Send(emailMsg)
	}

	func (h *UserHandler) generatePlainTextOTP(data templates.OTPEmailData) string {
		return `Hello ` + data.Name + `,
	
	Your verification code is: ` + data.OTP + `
	
	This code will expire in ` + strconv.Itoa(data.ExpiryMinutes) + ` minutes.
	
	If you didn't request this, please ignore this email.
	
	Best regards,
	The ` + data.AppName + ` Team`
	}