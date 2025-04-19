package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ProductName string `json:"product_name" bson:"product_name"`
	Price       int    `json:"price" bson:"price"`
	ProductCode string `json:"product_code" bson:"product_code"`
}

type Sale struct {
	SoldQuantity int    `json:"sold_quantity" bson:"sold_quantity"`
	Place        string `json:"place" bson:"place"`
	ProductCode  string `json:"product_code" bson:"product_code"`
}

type SaleSummary struct {
	Place       string `json:"place"`
	ProductName string `json:"product_name"`
	TotalSold   int    `json:"total_sold"`
}

var Mydb *gorm.DB

func Create_product(c *gin.Context) {

	var temp Product

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err := Mydb.Create(&temp).Error

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": temp})
}

func Entry_sale(c *gin.Context) {

	var temp Sale

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err := Mydb.Create(&temp).Error

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": temp})
}

func Sale_view(c *gin.Context) {

	var summaries []SaleSummary

	err := Mydb.Raw(`
	
	          select sales.place,    
			  products.product_name, 
			  sum(sales.sold_quantity) as total_sold 
			  from products
              inner join sales on  products.product_code = sales.product_code group by sales.place,
			  products.product_name;
	
	`).Scan(&summaries).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summaries)
}

func main() {

	dsn := "root:Vijay@123@tcp(localhost:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Mydb = db

	if err != nil {

		log.Fatal(err.Error())
	}

	db.AutoMigrate(&Sale{}, &Product{})

	router := gin.Default()

	router.POST("/product", Create_product)
	router.POST("/sale", Entry_sale)
	router.GET("/sale_view", Sale_view)
	router.Run(":7000")
}
