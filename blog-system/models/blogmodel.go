package models

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
