package models_test

import (
	"testing"
	"time"

	"go-echo-clean-architecture/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// UserModelTestSuite adalah suite pengujian untuk model User
type UserModelTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SetupSuite mempersiapkan koneksi database untuk pengujian
func (suite *UserModelTestSuite) SetupSuite() {
	// Gunakan database test terpisah untuk pengujian
	dsn := "host=localhost user=postgres password=admin123* dbname=db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error
	suite.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		suite.T().Fatalf("Gagal terhubung ke database test: %v", err)
	}

	// Migrasi skema database untuk pengujian
	err = suite.db.AutoMigrate(&models.User{})
	if err != nil {
		suite.T().Fatalf("Gagal melakukan migrasi database: %v", err)
	}
}

// TearDownSuite membersihkan database setelah pengujian selesai
func (suite *UserModelTestSuite) TearDownSuite() {
	// Hapus semua data user setelah pengujian
	suite.db.Exec("DELETE FROM users")

	// Tutup koneksi database
	sqlDB, err := suite.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

// SetupTest membersihkan database sebelum setiap test
func (suite *UserModelTestSuite) SetupTest() {
	// Hapus semua data user sebelum setiap test
	suite.db.Exec("DELETE FROM users")
}

// TestCreateUser menguji pembuatan user baru
func (suite *UserModelTestSuite) TestCreateUser() {
	// Buat user baru
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser",
		FullName: "Test User",
		Password: "password123",
		Email:    "test@example.com",
		Role:     "user",
		Phone:    "081234567890",
		Address:  "Jl. Test No. 123",
		Age:      25,
		Auditable: models.Auditable{
			CreatedAt: time.Now(),
			CreatedBy: uuid.New(),
			UpdatedAt: time.Now(),
			UpdatedBy: uuid.New(),
		},
	}

	// Simpan user ke database
	result := suite.db.Create(user)

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), result.Error)

	// Verifikasi user tersimpan di database
	var savedUser models.User
	suite.db.First(&savedUser, "id = ?", user.ID)
	assert.Equal(suite.T(), user.ID, savedUser.ID)
	assert.Equal(suite.T(), user.Username, savedUser.Username)
	assert.Equal(suite.T(), user.Email, savedUser.Email)
}

// TestUserRepositorySuite menjalankan suite pengujian
func TestUserModelSuite(t *testing.T) {
	suite.Run(t, new(UserModelTestSuite))
}
