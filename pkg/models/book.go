package models

import (
	"github.com/Aditya7feb/bookstore-golang/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllBooks() []Book {
	var Books []Book

	db.Find(&Books)

	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book

	db := db.Where("id = ?", id).First(&getBook)

	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book

	db.Where("ID = ?", id).Delete(book)

	return book
}
