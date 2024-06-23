package main

import (
	"RESTAPi/controller"
	"RESTAPi/middleware"
	"RESTAPi/model"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Vijay@123@tcp(localhost:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&model.Account{}, &model.User{})

	myserver := gin.Default()

	bank := model.Newmodel(db)
	branch := controller.Newcontroller(bank)

	// Public routes
	myserver.POST("/register", branch.Register)
	myserver.POST("/login", branch.Login)

	// Protected routes
	protected := myserver.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/NewAccount/", branch.Accountcreation)
		protected.GET("/Account/:id", branch.Getbyname)
		protected.GET("/Account/", branch.View)
		protected.PUT("/Account/:id", branch.Updateaccount)
		protected.DELETE("/Account/:id", branch.Deleteaccount)
	}

	myserver.Run("localhost:8080")
}
