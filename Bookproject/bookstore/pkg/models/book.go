package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	//gorm.Model         // `gorm:""json:name`// structure help us store in datbase
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	config.Connect() // help to connect database
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // based struct field db.AutoMigrate() automatically creating the table name book and column is struct filed
} // ORM  Object Relational mapping

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // is a function used to check if b is a new record in the database.if yes return true or false
	db.Create(&b)   // db.NewRecord(b) return true(new record) db.Create(&b)  insert in to the database.
	return b
}

func GetAllBooks() []Book {

	var Books []Book

	db.Find(&Books) // retrive  from the database which is select * from table and returning as a slice.
	return Books
}

func GetBookID(Id int64) (*Book, *gorm.DB) {

	var GetBook Book

	db := db.Where("ID=?", Id).Find(&GetBook)

	/*
		In the code db.Where("ID=?", Id).Find(&GetBook), you are querying the database to find a record where the ID column matches the value Id. Once found,
		the data corresponding to that record will be stored into the variable GetBook.
	*/

	return &GetBook, db
}

func DeleteBook(ID int64) Book {

	var book Book
	db.Where("ID=?", ID).Delete(book) // matching the id and delecting.

	return book
}

// ORM (Object-Relational Mapping) is nothing but by using gorm package will access the database crud operation.
// such is create and update delete in database without use querry.
