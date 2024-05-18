package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model  //this adds default ID, createdAt fields
	Name        string
	Publication string
	Author      string
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{}) //this is require to auto-increment ID
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var book []Book
	db.Find(&book)
	return book
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getMeBooK Book
	db := db.Where("ID=?", Id).Find(&getMeBooK)
	return &getMeBooK, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(&book)
	return book
}

func GetBookByIdOk(Id int64) (*Book, *gorm.DB) {
	var getMeMyBook Book
	db := db.Where("Id=?", Id).Find(&getMeMyBook)
	return &getMeMyBook, db

}
