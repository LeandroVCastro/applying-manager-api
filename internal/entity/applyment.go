package entity

import (
	"time"

	"gorm.io/gorm"
)

type Applyment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description *string        `json:"description"`
	Link        *string        `json:"link"`
	CompanyID   *int           `json:"company_id"`
	Company     Company        `json:"company"`
	PlatformId  *int           `json:"platform_id"`
	Platform    Platform       `json:"platform"`
	AppliedAt   time.Time      `json:"applied_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
