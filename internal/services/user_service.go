package services

import (
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/repositories"
	"go-echo-clean-architecture/internal/utils"
	"go-echo-clean-architecture/internal/validation"

	"github.com/google/uuid"
)

// UserService Step 1: Define the UserService interface
type UserService interface {
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetAllUsersWithPagination(pagination utils.PaginationType) ([]*models.User, int64, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id uuid.UUID, user *models.User) (*models.User, error)
	DeleteUser(id uuid.UUID, deleteBy uuid.UUID) error
}

// UserServiceImpl Step 2: Implement the UserService interface
type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

// NewUserServiceImpl UserRepository Step 3: Define the UserRepository interface
func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

// GetUserByID UserRepository Step 4: Implement the UserRepository interface
func (u *UserServiceImpl) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return nil, errors.NewNotFoundError("User not found", err)
	}
	return user, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to retrieve users", err)
	}
	return users, nil
}

func (u *UserServiceImpl) GetAllUsersWithPagination(pagination utils.PaginationType) ([]*models.User, int64, error) {
	users, totalRecords, err := u.userRepository.FindAllWithPagination(pagination)
	if err != nil {
		return nil, 0, errors.NewInternalServerError("Failed to retrieve users", err)
	}
	return users, totalRecords, nil
}

func (u *UserServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	// Validasi semua field wajib diisi
	err := validation.ValidateCreateUser(user)
	if err != nil {
		return nil, err
	}

	// Validasi existing user by email, username, and phone number
	if u.userRepository.ExistsByEmail(user.Email) {
		return nil, errors.NewBadRequestError("Email already exists", nil)
	}
	if u.userRepository.ExistsByUsername(user.Username) {
		return nil, errors.NewBadRequestError("Username already exists", nil)
	}
	if u.userRepository.ExistsByPhoneNumber(user.Phone) {
		return nil, errors.NewBadRequestError("Phone number already exists", nil)
	}

	// Generate UUID baru untuk user
	user.ID = uuid.New()

	// Generate password
	user.Password, err = utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate password")
	}

	createdUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to create user", err)
	}
	return createdUser, nil
}

func (u *UserServiceImpl) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	// Dapatkan user yang ada untuk dibandingkan
	existingUser, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Validasi dan hanya update field yang berubah
	updateData, err := validation.ValidateUpdateUser(existingUser, user)
	if err != nil {
		return nil, err
	}

	// Update user dengan data yang berubah saja
	updatedUser, err := u.userRepository.Update(id, updateData)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to update user", err)
	}
	return updatedUser, nil
}

func (u *UserServiceImpl) DeleteUser(id uuid.UUID, deleteBy uuid.UUID) error {
	// Validate user exists
	_, err := u.GetUserByID(id)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError("Failed to delete user", err)
	}
	return nil
}
