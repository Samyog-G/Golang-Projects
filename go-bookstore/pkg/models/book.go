package models

import (
	"github.com/Samyog-G/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()    //connect with connect func in config
	db = config.GetDB() //connect with GetDB func in config
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book { //(b*Book lets you receive the data in the type book and *Book returns te book)
	db.NewRecord(b) //received smth of type book
	db.Create(&b)   //created smth of type book
	return b        //returns type *Book
}

func GetBook() []Book { //using slice cause you want to return a book or a slice of book
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
