package config

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "akhil(username):password/tablenamecharset=utf8&parseTime=True&loc=Local")
	// Imp for sql: charset=utf8&parseTime=True&loc=Local"

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
