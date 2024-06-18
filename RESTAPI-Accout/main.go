package main

import (
	"RESTAPi/controller"
	"RESTAPi/model"
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "root:Vijay@123@tcp(localhost:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal(err.Error())
	}
	db.AutoMigrate(&model.Account{})

	myserver := gin.Default()

	bank := model.Newmodel(db)

	branch := controller.Newcontroller(bank)

	myserver.POST("/NewAccount/", branch.Accountcreation)
	myserver.GET("/Account/:id", branch.Getbyname)
	myserver.GET("/Account/", branch.View)
	myserver.PUT("/Account/:id", branch.Updateaccount)
	myserver.DELETE("/Account/:id", branch.Deleteaccount)

	myserver.Run("localhost:8080")

}
