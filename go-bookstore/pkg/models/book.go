package models

import{
	"github.com/Samyog-G/go-bookstore/pkg/config"

}

var db *gorm.db

type Book struct{
	gorm.model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect() //connect with connect func in config
	db=config.GetDB() //connect with GetDB func in config
	db.AutoMigrate(&Book{})
}

