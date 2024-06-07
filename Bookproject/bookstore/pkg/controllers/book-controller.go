package controllers

import (
	"bookstore/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {

	newBooks := models.GetAllBooks()

	c.JSON(http.StatusOK, newBooks)
}

func GetBookID(c *gin.Context) {

	bookID := c.Param("bookID")
	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	bookDetails, _ := models.GetBookID(ID)

	// Return book details in the response
	c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {

	Createbook := &models.Book{}

	//utils.Parsebody(c.Request, Createbook)
	c.ShouldBindJSON(Createbook)
	b := Createbook.CreateBook()
	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {

	bookID := c.Param("bookID")

	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	del := models.DeleteBook(ID)
	if err != nil {
		// Handle error (e.g., log error, return error response)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete details"})
		return
	}

	c.JSON(http.StatusOK, del)

}

func UpdateBook(c *gin.Context) {

	var updatebook = &models.Book{}

	//utils.Parsebody(c.Request, updatebook)
	c.ShouldBindJSON(updatebook) // Client sending data which is json format that we are converting to go format
	bookID := c.Param("bookID")
	ID, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing"})
		return
	}

	bookdetails, db := models.GetBookID(ID)

	if updatebook.Name != "" {

		bookdetails.Name = updatebook.Name

	}
	if updatebook.Author != "" {

		bookdetails.Author = updatebook.Author
	}

	if updatebook.Publication != "" {

		bookdetails.Publication = updatebook.Publication
	}
	db.Save(&bookdetails)

	c.JSON(http.StatusOK, bookdetails)
}
