package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Id       int
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Mobile   string `gorm:"uniqueIndex"`
	Password string
}
