package main

import (
	"log"
	"restapi/Model"
	"restapi/controller"

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

	db.AutoMigrate(&Model.Student{})

	repo := Model.NewRepository(db)

	ctrl := controller.Mycontroll(repo)

	r := gin.Default()

	r.GET("/student/", ctrl.Getstudent)
	r.POST("/student/", ctrl.Createstudent)
	r.GET("/student/:SID", ctrl.GetbyID)
	r.PUT("/student/:SID", ctrl.Updatestudent)
	r.DELETE("/student/:SID", ctrl.Delete)

	log.Fatal(r.Run("localhost:9010"))
}
