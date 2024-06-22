package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {

	app := gin.Default()
	api := app.Group("/api/v1")
	api.POST("/comments", CreateComment)
	app.Run(":3000")

}

func CreateComment(c *gin.Context) {
	var cmt Comment // Assuming Comment is a struct representing your comment data
	if err := c.BindJSON(&cmt); err != nil {
		// Log the error
		log.Println(err)

		// Return a JSON response with the error message and status code 400
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Your code to handle creating comments goes here using the 'cmt' struct
	// For example, you can insert the comment into a database

	// Return a JSON response indicating success
	PushCommentToqueue()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Comment created successfully",
	})
}
