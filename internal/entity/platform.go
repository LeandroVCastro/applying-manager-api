package entity

import (
	"time"

	"gorm.io/gorm"
)

type Platform struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Website   *string        `json:"website"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Applyments []Applyment `json:"applyments"`
}
