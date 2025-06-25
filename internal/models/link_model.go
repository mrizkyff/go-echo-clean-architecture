package models

import "github.com/google/uuid"

type Link struct {
	ID           uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	OriginalLink string    `json:"original_link" gorm:"column:original_link;type:varchar(255);not null" validate:"required"`
	ShortenLink  string    `json:"shorten_link" gorm:"column:shorten_link;type:varchar(255);not null" validate:"required"`
	UserID       uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`

	// Auditable
	Auditable
}
