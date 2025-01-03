package models

import (
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
}
