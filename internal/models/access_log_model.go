package models

import (
	"github.com/google/uuid"
	"time"
)

type AccessLog struct {
	ID         uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	AccessTime time.Time `json:"access_time" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	IpAddress  string    `json:"ip_address" gorm:"column:shorten_link;type:varchar(255);not null" validate:"required"`
	ClientInfo string    `json:"client_info" gorm:"column:client_info;type:varchar(255);not null" validate:"required"`
	LinkID     uuid.UUID `json:"link_id" gorm:"column:link_id;type:uuid" validate:"required"`
	UserID     uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid;not null"`

	Link Link `json:"link" gorm:"foreignKey:LinkID"`
	User User `json:"user" gorm:"foreignKey:UserID"`
}
