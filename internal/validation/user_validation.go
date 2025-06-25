package validation

import (
	"go-echo-clean-architecture/internal/errors"
	"go-echo-clean-architecture/internal/models"

	"github.com/go-playground/validator/v10"
)

// Inisialisasi validator
var validate = validator.New()

// ValidateCreateUser memvalidasi semua field wajib diisi saat create user menggunakan validator
func ValidateCreateUser(user *models.User) error {
	if user == nil {
		return errors.NewBadRequestError("User data is required")
	}

	// Gunakan validator untuk memvalidasi struct berdasarkan tag validate
	err := validate.Struct(user)
	if err != nil {
		// Jika terjadi error validasi, ambil error pertama
		if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
			firstError := validationErrors[0]
			switch firstError.Field() {
			case "Username":
				return errors.NewBadRequestError("Username is required")
			case "FullName":
				return errors.NewBadRequestError("Full name is required")
			case "Password":
				return errors.NewBadRequestError("Password is required")
			case "Email":
				if firstError.Tag() == "email" {
					return errors.NewBadRequestError("Invalid email format")
				}
				return errors.NewBadRequestError("Email is required")
			case "Role":
				return errors.NewBadRequestError("Role is required")
			case "Phone":
				return errors.NewBadRequestError("Phone is required")
			case "Address":
				return errors.NewBadRequestError("Address is required")
			case "Age":
				if firstError.Tag() == "gt" {
					return errors.NewBadRequestError("Age must be greater than 0")
				}
				return errors.NewBadRequestError("Age is required")
			default:
				return errors.NewBadRequestError(firstError.Error())
			}
		}
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

// ValidateUpdateUser memvalidasi field yang diupdate menggunakan validator
func ValidateUpdateUser(existingUser, updatedUser *models.User) (*models.User, error) {
	// Buat user baru yang hanya berisi field yang berubah
	result := &models.User{ID: existingUser.ID}

	// Bandingkan setiap field dan hanya update yang berubah
	if updatedUser.Username != "" && updatedUser.Username != existingUser.Username {
		result.Username = updatedUser.Username
		// Validasi username jika diubah
		if err := validate.Var(result.Username, "required"); err != nil {
			return nil, errors.NewBadRequestError("Username is required")
		}
	}

	if updatedUser.FullName != "" && updatedUser.FullName != existingUser.FullName {
		result.FullName = updatedUser.FullName
		// Validasi fullname jika diubah
		if err := validate.Var(result.FullName, "required"); err != nil {
			return nil, errors.NewBadRequestError("Full name is required")
		}
	}

	if updatedUser.Password != "" && updatedUser.Password != existingUser.Password {
		result.Password = updatedUser.Password
		// Validasi password jika diubah
		if err := validate.Var(result.Password, "required"); err != nil {
			return nil, errors.NewBadRequestError("Password is required")
		}
	}

	if updatedUser.Email != "" && updatedUser.Email != existingUser.Email {
		result.Email = updatedUser.Email
		// Validasi email jika diubah
		if err := validate.Var(result.Email, "required,email"); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
				if validationErrors[0].Tag() == "email" {
					return nil, errors.NewBadRequestError("Invalid email format")
				}
			}
			return nil, errors.NewBadRequestError("Email is required")
		}
	}

	if updatedUser.Role != "" && updatedUser.Role != existingUser.Role {
		result.Role = updatedUser.Role
		// Validasi role jika diubah
		if err := validate.Var(result.Role, "required"); err != nil {
			return nil, errors.NewBadRequestError("Role is required")
		}
	}

	if updatedUser.Phone != "" && updatedUser.Phone != existingUser.Phone {
		result.Phone = updatedUser.Phone
		// Validasi phone jika diubah
		if err := validate.Var(result.Phone, "required"); err != nil {
			return nil, errors.NewBadRequestError("Phone is required")
		}
	}

	if updatedUser.Address != "" && updatedUser.Address != existingUser.Address {
		result.Address = updatedUser.Address
		// Validasi address jika diubah
		if err := validate.Var(result.Address, "required"); err != nil {
			return nil, errors.NewBadRequestError("Address is required")
		}
	}

	if updatedUser.Age > 0 && updatedUser.Age != existingUser.Age {
		result.Age = updatedUser.Age
		// Validasi age jika diubah
		if err := validate.Var(result.Age, "gt=0"); err != nil {
			return nil, errors.NewBadRequestError("Age must be greater than 0")
		}
	}

	return result, nil
}
