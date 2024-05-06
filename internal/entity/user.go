package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string
}
