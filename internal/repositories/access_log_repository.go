package repositories

import (
	errors2 "errors"
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AccessLogRepository interface {
	Create(log *models.AccessLog) (*models.AccessLog, error)
	GetById(id uuid.UUID) (*models.AccessLog, error)
	FindAllWithPagination(pagination utils.PaginationType) ([]*models.AccessLog, int64, error)
}

type AccessLogRepositoryImpl struct {
	db *gorm.DB
}

func NewAccessLogRepositoryImpl(db *gorm.DB) *AccessLogRepositoryImpl {
	// Enable SQL query logging for all operations in this repository
	loggedDB := db.Session(&gorm.Session{
		Logger: db.Logger.LogMode(logger.Info),
	})
	return &AccessLogRepositoryImpl{db: loggedDB}
}

func (a *AccessLogRepositoryImpl) Create(log *models.AccessLog) (*models.AccessLog, error) {
	if log.ID != uuid.Nil {
		log.ID = uuid.New()
	}

	err := a.db.Create(log).Error
	if err != nil {
		return nil, errors.NewInternalServerError("Gagal membuat access log")
	}

	return log, nil
}

func (a *AccessLogRepositoryImpl) GetById(id uuid.UUID) (*models.AccessLog, error) {
	existingLog := models.AccessLog{}

	err := a.db.Where(models.AccessLog{ID: id}).First(&existingLog).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewInternalServerError("Access log tidak ditemukan")

		}
		return nil, errors.NewInternalServerError("Gagal mendapatkan access log")
	}

	return &existingLog, nil
}

func (a *AccessLogRepositoryImpl) FindAllWithPagination(pagination utils.PaginationType) ([]*models.AccessLog, int64, error) {
	var logs []*models.AccessLog
	var totalElements int64

	// Count all - explicitly debug this query for performance monitoring
	countQuery := a.db.Model(models.AccessLog{})
	err := countQuery.Count(&totalElements).Error
	if err != nil {
		return nil, 0, errors.NewInternalServerError("Gagal counting access log")
	}

	// Calculate offset
	offset := (pagination.Page - 1) * pagination.PageSize

	// Build query with sorting
	query := a.db.Preload("User").Preload("Link")

	// Add sorting
	if pagination.SortBy == "asc" {
		query = query.Order(pagination.SortBy + " asc")
	} else {
		query = query.Order(pagination.SortBy + " desc")
	}

	// Execute final query
	err = query.Limit(pagination.PageSize).Offset(offset).Find(&logs).Error
	if err != nil {
		return nil, 0, errors.NewInternalServerError("Gagal mendapatkan data access log")
	}

	return logs, totalElements, nil
}
