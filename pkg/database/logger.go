package database

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConfigureGormLogger sets up SQL query logging for GORM
func ConfigureGormLogger(logLevel logger.LogLevel) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold (1 second)
			LogLevel:                  logLevel,    // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable color
		},
	)
}

// EnableQueryLogging configures a GORM DB instance to log SQL queries
func EnableQueryLogging(db *gorm.DB) *gorm.DB {
	logrus.Info("Enabling SQL query logging for GORM")
	return db.Session(&gorm.Session{
		Logger: ConfigureGormLogger(logger.Info),
	})
}

// EnableErrorLogging configures a GORM DB instance to log only errors
func EnableErrorLogging(db *gorm.DB) *gorm.DB {
	logrus.Info("Enabling SQL error logging for GORM")
	return db.Session(&gorm.Session{
		Logger: ConfigureGormLogger(logger.Error),
	})
}
