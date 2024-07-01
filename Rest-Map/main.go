package main

import (
	"log"
	"restapi/Model"
	"restapi/controller"
	"restapi/middleware"

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

	db.AutoMigrate(&Model.Student{}, &Model.Register{})

	repo := Model.NewRepository(db)

	ctrl := controller.Mycontroll(repo)

	r := gin.Default()
	r.POST("/Newuser", ctrl.NewUser)
	r.POST("/Login", ctrl.Mylogin)

	protect := r.Group("/")

	protect.Use(middleware.AuthMiddleware())
	{
		protect.GET("/student/", ctrl.Getstudent)
		protect.POST("/student/", ctrl.Createstudent)
		protect.GET("/student/:SID", ctrl.GetbyID)
		protect.PUT("/student/:SID", ctrl.Updatestudent)
		protect.DELETE("/student/:SID", ctrl.Delete)
	}
	//r.Run("localhost:9010")
	r.Run("0.0.0.0:9010")

}

// "email": "test@gmail.com",
// "password": "Vijay@123"

/*
{
 user 1.

    "name" : "Check",
    "email": "check@gmail.com",
    "password": "check@1235"
}
user 2.

{
     "name": "Johan",
    "email": "johan@gmail.com",
    "password": "johan@123"
}

   input format
{

  "name": "check",
  "place": "city",
  "contactnumber": "1234567890",
  "dob": "2000-01-01",
  "user_id": 101
}

    "name" : "Imay",
    "email": "imay@gmail.com",
    "password": "Imay@123"

*/
