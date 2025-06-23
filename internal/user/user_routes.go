package user

import (
	"database/sql"

	"github.com/Dev-Awaab/go_job_portal_api/config"
	"github.com/Dev-Awaab/go_job_portal_api/internal/otp"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/mail"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, dbConn *sql.DB, cfg *config.Config,mailer mail.Mailer) {
	otpRepo := otp.NewOtpRepository(dbConn)
	otpSvc := otp.NewOtpService(otpRepo)
	// Initialize dependencies
	repo := NewUserRepository(dbConn)
	service := NewUserService(repo)
	handler := NewUserHandler(service, mailer, otpSvc)

	// Public routes
	r.POST("/register", handler.Register)
	r.POST("/verify-email", handler.VerifyUserEmail)
	r.POST("/login", handler.Login)
	// r.POST("/login", handler.Login)

	// Authenticated routes (require JWT)
	// authRoutes := r.Group("")
	// authRoutes.Use(middleware.AuthMiddleware())
	// {
	// 	authRoutes.GET("/profile", handler.GetProfile)
	// 	authRoutes.PUT("/profile", handler.UpdateProfile)
	// 	authRoutes.PUT("/password", handler.ChangePassword)
	// 	authRoutes.DELETE("", handler.DeleteAccount)
	// }

	// // Admin routes
	// adminRoutes := r.Group("")
	// adminRoutes.Use(middleware.AdminMiddleware())
	// {
	// 	adminRoutes.GET("", handler.ListUsers) // Admin-only user listing
	// }
}