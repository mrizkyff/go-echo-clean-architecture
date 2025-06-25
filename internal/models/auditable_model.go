package models

import (
	"github.com/google/uuid"
	"time"
)

type Auditable struct {
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"column:created_by;type:uuid"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedBy uuid.UUID  `json:"updated_by" gorm:"column:updated_by;type:uuid"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:timestamp;default:NULL"`
	DeletedBy uuid.UUID  `json:"deleted_by" gorm:"column:deleted_by;type:uuid"`
}
