package routes

import (
	"bookstore/pkg/controllers"

	"github.com/gin-gonic/gin"
	// need controllers
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {

	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:bookID", controllers.GetBookID)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)
}

//http://localhost:9010/book/
//http://localhost:9010/book/2
