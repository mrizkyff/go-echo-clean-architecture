package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	Username string    `json:"username" gorm:"column:username;type:varchar(255);unique;not null" validate:"required"`
	FullName string    `json:"full_name" gorm:"column:full_name;type:varchar(255);not null" validate:"required"`
	Password string    `json:"password" gorm:"column:password;type:varchar(255);not null" validate:"required"`
	Email    string    `json:"email" gorm:"column:email;type:varchar(255);unique;not null" validate:"required,email"`
	Role     string    `json:"role" gorm:"column:role;type:varchar(255);not null" validate:"required"`
	Phone    string    `json:"phone" gorm:"column:phone;type:varchar(255);not null" validate:"required"`
	Address  string    `json:"address" gorm:"column:address;type:varchar(255);not null" validate:"required"`
	Age      int       `json:"age" gorm:"column:age;type:int;not null" validate:"required,gt=0"`
	Links    []Link    `json:"links" gorm:"foreignKey:UserID"`

	// Auditable
	Auditable
}
