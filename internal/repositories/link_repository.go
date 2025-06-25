package repositories

import (
	"errors"
	errors2 "go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LinkRepository interface {
	FindById(id uuid.UUID) (*models.Link, error)
	FindAllWithPagination(pagination utils.PaginationType) ([]*models.Link, int64, error)
	Create(link *models.Link) (*models.Link, error)
	Update(id uuid.UUID, link *models.Link) (*models.Link, error)
	Delete(id uuid.UUID) error
	ExistsByOriginalLink(originalLink string) (bool, error)
	ExistsByShortenLink(shortenLink string) (bool, error)
}

type LinkRepositoryImpl struct {
	db *gorm.DB
}

func NewLinkRepositoryImpl(db *gorm.DB) *LinkRepositoryImpl {
	return &LinkRepositoryImpl{db: db}
}

func (l *LinkRepositoryImpl) FindById(id uuid.UUID) (*models.Link, error) {
	link := models.Link{}

	// Find
	result := l.db.Preload("User").Where(&models.Link{ID: id}).First(&link)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError("Link not found", result.Error)
		}

		return nil, errors2.NewInternalServerError("Internal server error", result.Error)
	}
	return &link, nil
}

func (l *LinkRepositoryImpl) FindAllWithPagination(pagination utils.PaginationType) ([]*models.Link, int64, error) {
	var links []*models.Link
	var totalRecords int64

	// Count total records
	err := l.db.Model(&models.Link{}).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (pagination.Page - 1) * pagination.PageSize

	// Build query with sorting
	query := l.db.Preload("User")

	// Add sorting
	if pagination.SortDir == "asc" {
		query = query.Order(pagination.SortBy + " asc")
	} else {
		query = query.Order(pagination.SortBy + " desc")
	}

	err = query.Limit(pagination.PageSize).Offset(offset).Find(&links).Error
	if err != nil {
		return nil, 0, err
	}

	return links, totalRecords, nil
}

func (l *LinkRepositoryImpl) Create(link *models.Link) (*models.Link, error) {
	if link.ID == uuid.Nil {
		link.ID = uuid.New()
	}

	err := l.db.Create(link).Error
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (l *LinkRepositoryImpl) Update(id uuid.UUID, link *models.Link) (*models.Link, error) {
	existingLink := models.Link{}

	// Find existing link
	err := l.db.Where(&models.Link{ID: id}).First(&existingLink).Error
	if err != nil {
		return nil, err
	}

	// Fill different field
	if existingLink.OriginalLink != link.OriginalLink {
		existingLink.OriginalLink = link.OriginalLink
	}

	if existingLink.ShortenLink != link.ShortenLink {
		existingLink.ShortenLink = link.ShortenLink
	}

	// Fill auditable
	existingLink.UpdatedAt = time.Now()

	err = l.db.Save(&existingLink).Error
	if err != nil {
		return nil, err
	}

	return &existingLink, nil
}

func (l *LinkRepositoryImpl) Delete(id uuid.UUID) error {
	link := models.Link{}

	result := l.db.Where(&models.Link{ID: id}).First(&link)
	if result != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors2.NewNotFoundError("Link not found", result.Error)
		}
	}

	deleteError := l.db.Delete(&link).Error
	if deleteError != nil {
		return deleteError
	}

	return nil
}

func (l *LinkRepositoryImpl) ExistsByOriginalLink(originalLink string) (bool, error) {
	link := models.Link{}

	result := l.db.Where(&models.Link{OriginalLink: originalLink}).First(&link)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (l *LinkRepositoryImpl) ExistsByShortenLink(shortenLink string) (bool, error) {
	link := models.Link{}

	result := l.db.Where(&models.Link{ShortenLink: shortenLink}).First(&link)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}
