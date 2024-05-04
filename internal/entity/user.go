package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Password string
	// CreatedAt time.Time
	// UpdatedAt time.Time
}
