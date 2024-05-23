package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description *string        `json:"description"`
	Website     *string        `json:"website"`
	Linkedin    *string        `json:"linkedin"`
	Glassdoor   *string        `json:"glasdoor"`
	Instagram   *string        `json:"instagram"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Applyments []Applyment `json:"applyments"`
}
