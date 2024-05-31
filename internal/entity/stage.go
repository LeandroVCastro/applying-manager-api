package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stage struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"not null" json:"description"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Applyments []Applyment `json:"applyments,omitempty"`
}
