package main

import (
	"fmt"

	"github.com/Dev-Awaab/go_job_portal_api/config"
	"github.com/Dev-Awaab/go_job_portal_api/db"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/logger"
	"github.com/Dev-Awaab/go_job_portal_api/pkg/router"
)

func main() {
	// Initialize the logger
	_, err := logger.InitLogger(logger.InfoLevel, "stdout")
	if err != nil {
	 panic(err)
	}


	// Load Config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Error("Failed to load config file: %v", err)
	}

	fmt.Println("DBSource", cfg.DBSource)
	// Initalize DB
	dbConn := db.InitDB(cfg.DBSource)

	defer dbConn.Close()

	r := router.SetupRoutes(dbConn, &cfg)

	logger.Info("Server running on port %s", cfg.ServerPort)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		logger.Error("Failed to start server: %v", err)
		panic(err)
	}
	
}