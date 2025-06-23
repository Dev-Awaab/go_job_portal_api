package db

import (
	"database/sql"
	"time"

	"github.com/Dev-Awaab/go_job_portal_api/pkg/logger"
	_ "github.com/lib/pq"
)

func InitDB(dbSource string) *sql.DB {
  db, err := sql.Open("postgres", dbSource)

    // Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

  if err != nil {
	 logger.Error("Failed to connect to database: %v", err)
  }

	
  if err := db.Ping(); err != nil {
	logger.Error("Database connection error: %v", err)
}
logger.Info("Connected to the database successfully!")
return db
}