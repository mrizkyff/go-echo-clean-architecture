package database

import (
	"fmt"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/models/config"
	"time"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	DB     *gorm.DB
	config config.PostgresConfig
}

func NewPostgresClient(DB *gorm.DB, config config.PostgresConfig) *PostgresClient {
	return &PostgresClient{DB: DB, config: config}
}

func (receiver *PostgresClient) InitDBConnection() error {
	// Konfigurasi koneksi PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", receiver.config.Host, receiver.config.Username, receiver.config.Password, receiver.config.DBName, receiver.config.Port, receiver.config.SSLMode)

	var err error
	logrus.Info("*** Connecting to postgres ***")
	receiver.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gagal terhubung ke database: %v", err)
	}
	logrus.Info("*** Success connected to postgres ***")

	// Konfigurasi connection pooling
	sqlDB, err := receiver.DB.DB()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan sql.DB: %v", err)
	}

	// Konfigurasi connection pooling
	sqlDB.SetMaxIdleConns(10)           // Jumlah koneksi idle maksimum
	sqlDB.SetMaxOpenConns(100)          // Jumlah koneksi terbuka maksimum
	sqlDB.SetConnMaxLifetime(time.Hour) // Waktu hidup koneksi maksimum

	// Run auto migrations
	// Auto migrate model User
	err = receiver.DB.AutoMigrate(&models.User{})
	err = receiver.DB.AutoMigrate(&models.Link{})
	err = receiver.DB.AutoMigrate(&models.AccessLog{})
	if err != nil {
		return fmt.Errorf("gagal melakukan migrasi model: %v", err)
	}

	return nil
}

func (receiver *PostgresClient) GetDB() *gorm.DB {
	return receiver.DB
}

// CloseDB menutup koneksi database
func (receiver *PostgresClient) CloseDB() error {
	sqlDB, err := receiver.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
