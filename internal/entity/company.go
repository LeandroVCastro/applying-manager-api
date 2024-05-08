package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description *string
	Website     *string
	Linkedin    *string
	Glassdoor   *string
	Instagram   *string
}
