package repositories_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// UserRepositoryTestSuite adalah suite pengujian untuk repository User
type UserRepositoryTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository repositories.UserRepository
}

// SetupSuite mempersiapkan koneksi database untuk pengujian
func (suite *UserRepositoryTestSuite) SetupSuite() {
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

	// Inisialisasi repository
	suite.repository = repositories.NewUserRepositoryImpl(suite.db)
}

// TearDownSuite membersihkan database setelah pengujian selesai
func (suite *UserRepositoryTestSuite) TearDownSuite() {
	// Hapus semua data user setelah pengujian
	suite.db.Exec("DELETE FROM users")

	// Tutup koneksi database
	sqlDB, err := suite.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

// SetupTest membersihkan database sebelum setiap test
func (suite *UserRepositoryTestSuite) SetupTest() {
	// Hapus semua data user sebelum setiap test
	suite.db.Exec("DELETE FROM users")
}

// createTestUser membuat user untuk pengujian
func (suite *UserRepositoryTestSuite) createTestUser() *models.User {
	user := &models.User{
		ID:       uuid.New(),
		Username: "testuser" + uuid.New().String(),
		FullName: "Test User" + uuid.New().String(),
		Password: "password123",
		Email:    uuid.NewString() + "test@example.com",
		Role:     "user",
		Phone:    fmt.Sprintf("0812345678%02d", rand.Intn(100)),
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
	suite.db.Create(user)
	return user
}

// TestFindById menguji fungsi FindById
func (suite *UserRepositoryTestSuite) TestFindById() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Panggil fungsi FindById
	foundUser, err := suite.repository.FindById(user.ID)

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi user ditemukan
	assert.NotNil(suite.T(), foundUser)
	assert.Equal(suite.T(), user.ID, foundUser.ID)
	assert.Equal(suite.T(), user.Username, foundUser.Username)
	assert.Equal(suite.T(), user.Email, foundUser.Email)

	// Coba cari user dengan ID yang tidak ada
	nonExistentID := uuid.New()
	nonExistentUser, err := suite.repository.FindById(nonExistentID)

	// Verifikasi error
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), nonExistentUser)
}

// TestFindAll menguji fungsi FindAll
func (suite *UserRepositoryTestSuite) TestFindAll() {
	// Buat beberapa user untuk pengujian
	suite.createTestUser()
	suite.createTestUser() // Buat user kedua dengan data yang sama (ID akan berbeda)

	// Panggil fungsi FindAll
	users, err := suite.repository.FindAll()

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi jumlah user
	assert.Equal(suite.T(), 2, len(users))
}

// TestFindAllWithPagination menguji fungsi FindAllWithPagination
func (suite *UserRepositoryTestSuite) TestFindAllWithPagination() {
	// Buat beberapa user untuk pengujian
	for i := 0; i < 10; i++ {
		suite.createTestUser() // Buat 10 user dengan data yang sama (ID akan berbeda)
	}

	// Panggil fungsi FindAllWithPagination untuk halaman pertama
	users, total, err := suite.repository.FindAllWithPagination(*utils.NewPaginationType())

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi jumlah total user
	assert.Equal(suite.T(), int64(10), total)
	// Verifikasi jumlah user pada halaman pertama
	assert.Equal(suite.T(), 5, len(users))

	// Panggil fungsi FindAllWithPagination untuk halaman kedua
	users2, total2, err := suite.repository.FindAllWithPagination(*utils.NewPaginationType())

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi jumlah total user
	assert.Equal(suite.T(), int64(10), total2)
	// Verifikasi jumlah user pada halaman kedua
	assert.Equal(suite.T(), 5, len(users2))

	// Verifikasi user pada halaman pertama dan kedua berbeda
	assert.NotEqual(suite.T(), users[0].ID, users2[0].ID)
}

// TestCreate menguji fungsi Create
func (suite *UserRepositoryTestSuite) TestCreate() {
	// Buat user baru
	user := &models.User{
		Username: "newuser",
		FullName: "New User",
		Password: "password123",
		Email:    "new@example.com",
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

	// Panggil fungsi Create
	createdUser, err := suite.repository.Create(user)

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi user berhasil dibuat
	assert.NotNil(suite.T(), createdUser)
	assert.NotEqual(suite.T(), uuid.Nil, createdUser.ID) // ID seharusnya di-generate
	assert.Equal(suite.T(), user.Username, createdUser.Username)
	assert.Equal(suite.T(), user.Email, createdUser.Email)

	// Verifikasi user tersimpan di database
	var savedUser models.User
	suite.db.First(&savedUser, "id = ?", createdUser.ID)
	assert.Equal(suite.T(), createdUser.ID, savedUser.ID)
	assert.Equal(suite.T(), user.Username, savedUser.Username)
}

// TestUpdate menguji fungsi Update
func (suite *UserRepositoryTestSuite) TestUpdate() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Buat data update
	updatedUser := &models.User{
		FullName: "Updated User",
		Email:    "updated@example.com",
		Auditable: models.Auditable{
			UpdatedAt: time.Now(),
			UpdatedBy: uuid.New(),
		},
	}

	// Panggil fungsi Update
	resultUser, err := suite.repository.Update(user.ID, updatedUser)

	fmt.Printf("user: %+v\n", user)
	fmt.Printf("updatedUser: %+v\n", updatedUser)

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)
	// Verifikasi user berhasil diupdate
	assert.NotNil(suite.T(), resultUser)
	assert.NotNil(suite.T(), resultUser.ID)
	assert.Equal(suite.T(), updatedUser.FullName, resultUser.FullName)
	assert.Equal(suite.T(), updatedUser.Email, resultUser.Email)

	// Verifikasi user terupdate di database
	var savedUser models.User
	suite.db.First(&savedUser, "id = ?", user.ID)
	assert.Equal(suite.T(), updatedUser.FullName, savedUser.FullName)
	assert.Equal(suite.T(), updatedUser.Email, savedUser.Email)
}

// TestDelete menguji fungsi Delete
func (suite *UserRepositoryTestSuite) TestDelete() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Panggil fungsi Delete
	err := suite.repository.Delete(user.ID)

	// Verifikasi tidak ada error
	assert.NoError(suite.T(), err)

	// Verifikasi user terhapus dari database
	var deletedUser models.User
	result := suite.db.First(&deletedUser, "id = ?", user.ID)
	assert.Error(suite.T(), result.Error)
	assert.Equal(suite.T(), "record not found", result.Error.Error())
}

// TestExistsByEmail menguji fungsi ExistsByEmail
func (suite *UserRepositoryTestSuite) TestExistsByEmail() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Panggil fungsi ExistsByEmail dengan email yang ada
	exists := suite.repository.ExistsByEmail(user.Email)

	// Verifikasi email ditemukan
	assert.True(suite.T(), exists)

	// Panggil fungsi ExistsByEmail dengan email yang tidak ada
	exists = suite.repository.ExistsByEmail("nonexistent@example.com")

	// Verifikasi email tidak ditemukan
	assert.False(suite.T(), exists)
}

// TestExistsByUsername menguji fungsi ExistsByUsername
func (suite *UserRepositoryTestSuite) TestExistsByUsername() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Panggil fungsi ExistsByUsername dengan username yang ada
	exists := suite.repository.ExistsByUsername(user.Username)

	// Verifikasi username ditemukan
	assert.True(suite.T(), exists)

	// Panggil fungsi ExistsByUsername dengan username yang tidak ada
	exists = suite.repository.ExistsByUsername("nonexistent")

	// Verifikasi username tidak ditemukan
	assert.False(suite.T(), exists)
}

// TestExistsByPhoneNumber menguji fungsi ExistsByPhoneNumber
func (suite *UserRepositoryTestSuite) TestExistsByPhoneNumber() {
	// Buat user untuk pengujian
	user := suite.createTestUser()

	// Panggil fungsi ExistsByPhoneNumber dengan nomor telepon yang ada
	exists := suite.repository.ExistsByPhoneNumber(user.Phone)

	// Verifikasi nomor telepon ditemukan
	assert.True(suite.T(), exists)

	// Panggil fungsi ExistsByPhoneNumber dengan nomor telepon yang tidak ada
	exists = suite.repository.ExistsByPhoneNumber("087654321098")

	// Verifikasi nomor telepon tidak ditemukan
	assert.False(suite.T(), exists)
}

// TestUserRepositorySuite menjalankan suite pengujian
func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
