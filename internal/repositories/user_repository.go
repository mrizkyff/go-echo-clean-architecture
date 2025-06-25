package repositories

import (
	"context"
	"errors"
	"fmt"
	errors2 "go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uuid.UUID) (*models.User, error)
	FindByUsername(username string, ctx context.Context) (*models.User, error)
	FindByEmail(email string, ctx context.Context) (*models.User, error)
	FindByPhone(phone string, ctx context.Context) (*models.User, error)
	FindAll() ([]*models.User, error)
	FindAllWithPagination(pagination utils.PaginationType) ([]*models.User, int64, error)
	Create(user *models.User) (*models.User, error)
	Update(id uuid.UUID, user *models.User) (*models.User, error)
	Delete(id uuid.UUID) error
	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
	ExistsByPhoneNumber(phoneNumber string) bool
	FindByUserNameOrEmailOrPhone(query string, ctx context.Context) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) FindByUserNameOrEmailOrPhone(query string, ctx context.Context) (*models.User, error) {
	var user = &models.User{}

	err := u.db.WithContext(ctx).Where("username = ? OR email = ? OR phone = ?", query, query, query).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError(fmt.Sprintf("User with username %s not found", query))
		}
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByUsername(username string, ctx context.Context) (*models.User, error) {
	var user = &models.User{}
	err := u.db.WithContext(ctx).Where(models.User{Username: username}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError(fmt.Sprintf("User with username %s not found", username))
		}
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string, ctx context.Context) (*models.User, error) {
	var user = &models.User{}
	err := u.db.WithContext(ctx).Where(models.User{Email: email}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError(fmt.Sprintf("User with email %s not found", email))
		}
	}

	return user, nil

}

func (u *UserRepositoryImpl) FindByPhone(phone string, ctx context.Context) (*models.User, error) {
	var user = &models.User{}
	err := u.db.WithContext(ctx).Where(models.User{Phone: phone}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors2.NewNotFoundError(fmt.Sprintf("User with phone %s not found", phone))
		}
	}

	return user, nil

}

func (u *UserRepositoryImpl) FindById(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := u.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindAllWithPagination(pagination utils.PaginationType) ([]*models.User, int64, error) {
	var users []*models.User
	var totalRecords int64

	// Count total records
	err := u.db.Model(&models.User{}).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (pagination.Page - 1) * pagination.PageSize

	// Build query with sorting
	query := u.db

	// Add sorting
	if pagination.SortDir == "asc" {
		query = query.Order(pagination.SortBy + " asc")
	} else {
		query = query.Order(pagination.SortBy + " desc")
	}

	// Get paginated records
	err = query.Limit(pagination.PageSize).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, totalRecords, nil
}

func (u *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	// Generate UUID jika belum ada
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Update(id uuid.UUID, user *models.User) (*models.User, error) {
	err := u.db.Model(user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) Delete(id uuid.UUID) error {
	err := u.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) ExistsByEmail(email string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("email =?", email).Count(&count)
	return count > 0
}

func (u *UserRepositoryImpl) ExistsByUsername(username string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("username =?", username).Count(&count)
	return count > 0
}

func (u *UserRepositoryImpl) ExistsByPhoneNumber(phoneNumber string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("phone =?", phoneNumber).Count(&count)
	return count > 0
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
