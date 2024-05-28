package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description *string        `json:"description,omitempty"`
	Website     *string        `json:"website,omitempty"`
	Linkedin    *string        `json:"linkedin,omitempty"`
	Glassdoor   *string        `json:"glasdoor,omitempty"`
	Instagram   *string        `json:"instagram,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Applyments []Applyment `json:"applyments,omitempty"`
}
