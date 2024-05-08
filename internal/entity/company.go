package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID          uint
	Name        string `gorm:"not null"`
	Description *string
	Website     *string
	Linkedin    *string
	Glassdoor   *string
	Instagram   *string
}
