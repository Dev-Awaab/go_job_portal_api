package router

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Dev-Awaab/go_job_portal_api/config"
	"github.com/Dev-Awaab/go_job_portal_api/internal/user"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/mail"
	"github.com/gin-gonic/gin"
)



func SetupRoutes(dbConn *sql.DB, cfg *config.Config)*gin.Engine  {
	r := gin.Default()
	
	// Global middleware
	// r.Use(middleware.CORS())
	// r.Use(middleware.Logger())

	// Initialize mailer (provider choice could come from env var)
	mailer, err := mail.NewMailer(mail.Resend, *cfg)
	if err != nil {
		log.Fatalf("Failed to create mailer: %v", err)
	}

	
	
	// API base route
	api := r.Group("/api/v1")
	{
			// Health check
			api.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{"status": "ok"})
			})
		// User routes
		userGroup := api.Group("/users")
		user.SetupUserRoutes(userGroup, dbConn, cfg, mailer)
		
		// Job routes
		jobGroup := api.Group("/jobs")
		fmt.Println(jobGroup)
		// job.SetupJobRoutes(jobGroup, dbConn, cfg)
		
	
	}
	
 return r
}

