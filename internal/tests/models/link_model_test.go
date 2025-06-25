package models_test

import (
	"encoding/json"
	"fmt"
	"go-echo-clean-architecture/internal/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LinkModelTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *LinkModelTestSuite) SetupSuite() {
	// Gunakan database test terpisah untuk pengujian
	dsn := "host=localhost user=postgres password=admin123* dbname=db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error
	suite.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		suite.T().Fatalf("Gagal terhubung ke database test: %v", err)
	}

	// Migrasi skema database untuk pengujian
	err = suite.db.AutoMigrate(&models.User{})
	err = suite.db.AutoMigrate(&models.Link{})
	if err != nil {
		suite.T().Fatalf("Gagal melakukan migrasi database: %v", err)
	}
}

func (suite *LinkModelTestSuite) TestCreate() {
	link := models.Link{
		ID:           uuid.New(),
		OriginalLink: "https://example.com/original2",
		ShortenLink:  "https://go-echo-clean-architecture/shrt3",
		Auditable: models.Auditable{
			CreatedAt: time.Time{},
			CreatedBy: uuid.New(),
			UpdatedAt: time.Time{},
			UpdatedBy: uuid.New(),
			DeletedAt: nil,
			DeletedBy: uuid.New(),
		},
	}

	err := suite.db.Create(&link).Error
	if err != nil {
		logrus.Errorf("Error %s", err)
	}
	logrus.Info(link)
}

func (suite *LinkModelTestSuite) TestFindAll() {
	var links []models.Link

	err := suite.db.Find(&links).Error
	if err != nil {
		logrus.Errorf("Error %s", err)
	}

	logrus.Info(links)

	for _, link := range links {
		marshal, err := json.MarshalIndent(link, "", "    ")
		if err != nil {
			logrus.Errorf("Error: %s", err)
		}

		fmt.Print(string(marshal))
	}
}

// Menambahkan fungsi TestMain untuk menjalankan suite test
func TestLinkModelSuite(t *testing.T) {
	suite.Run(t, new(LinkModelTestSuite))
}
