package main

import (
	"log"
	model "restaurant_mysql_db/Model"
	"restaurant_mysql_db/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	dsn := "root:Vijay@123@tcp(localhost:3306)/myhotel?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal(err.Error())
	}

	db.AutoMigrate(&model.Food{}, &model.Order_item{}, &model.Table{}, &model.Invoice{})

	myserver := gin.Default()

	myserver.Use(func(ctx *gin.Context) {

		ctx.Set("db", db)
		ctx.Next()
	})

	foodDb := model.Food_repo(db)

	Order_itemdb := model.Order_repo(db)
	TableDB := model.Table_Repo(db)
	Invoice := model.Invoice_Repo(db)

	Invoiceservice := controller.Invoice_controller(Invoice)
	foodservice := controller.Food_controll(foodDb)
	orderservice := controller.OrderItem_controller(Order_itemdb)
	tableservice := controller.Table_controller(TableDB)

	myserver.GET("/table", tableservice.ViewTables)
	myserver.GET("/order/:table_number", orderservice.ViewyTable)
	myserver.GET("/order", orderservice.Getall)
	myserver.GET("/food", foodservice.Getfoods)

	myserver.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	myserver.POST("/food", foodservice.Creatfood)
	myserver.POST("/table", tableservice.Maketable)
	myserver.POST("/order", orderservice.Creater_orderItem)
	myserver.POST("/invoice", Invoiceservice.Createbill)

	myserver.DELETE("/food/:foodid", foodservice.Remove_food)
	myserver.DELETE("/table/:table_id", tableservice.Remove_table)
	myserver.DELETE("/order/:order_item_id", orderservice.Remove_order_item)

	myserver.PUT("/table", tableservice.Update_table)
	myserver.PUT("/order", orderservice.Update_orderItem)

	myserver.Run("localhost:8000") // 192.168.1.7:
}
