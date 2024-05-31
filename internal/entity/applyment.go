package entity

import (
	"time"

	"gorm.io/gorm"
)

type Applyment struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description *string        `json:"description,omitempty"`
	Link        *string        `json:"link,omitempty"`
	CompanyId   *uint          `json:"company_id,omitempty"`
	Company     *Company       `json:"company,omitempty"`
	PlatformId  *uint          `json:"platform_id,omitempty" gorm:"default:null"`
	Platform    *Platform      `json:"platform,omitempty"`
	StageId     *uint          `json:"stage_id,omitempty" gorm:"default:null"`
	Stage       *Stage         `json:"stage,omitempty"`
	AppliedAt   *time.Time     `json:"applied_at,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
