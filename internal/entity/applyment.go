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
	CompanyID   *int           `json:"companyId"`
	Company     Company        `json:"company"`
	PlatformId  *int           `json:"platformId"`
	Platform    Platform       `json:"platform"`
	AppliedAt   time.Time      `json:"applieddAt"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
