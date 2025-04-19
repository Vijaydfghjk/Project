package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var productsCollection *mongo.Collection
var Salecollection *mongo.Collection

func Create_product(c *gin.Context) {

	var temp Product

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	_, err := productsCollection.InsertOne(context.TODO(), temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Successfully added"})
}

func Entry_sale(c *gin.Context) {

	var temp Sale

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	_, err := Salecollection.InsertOne(context.TODO(), temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to insert product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "successfully added"})
}

func Agggregate_sale(c *gin.Context) {

	pipline := mongo.Pipeline{

		{

			{"$group", bson.D{

				{"_id", bson.D{

					{"product_code", "$product_code"},
					{"place", "$place"},
				}},

				{"Total_sale", bson.D{

					{"$sum", "$sold_quantity"},
				}},
			}},
		},

		{

			{"$lookup", bson.D{

				{"from", "products"},
				{"localField", "_id.product_code"},
				{"foreignField", "product_code"},
				{"as", "Product_details"},
			}},
		},

		{

			{"$unwind", "$Product_details"},
		},

		{

			{"$project", bson.D{

				{"_id", 0},
				{"Product", "$Product_details.product_name"},
				{"place", "$_id.place"},
				{"Total_sale", 1},
			}},
		},
	}

	cursor, err := Salecollection.Aggregate(context.TODO(), pipline)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to aggregate sales"})
		return
	}

	defer cursor.Close(context.TODO())

	var result []bson.M

	if err := cursor.All(context.TODO(), &result); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to process results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Sale_data": result})
}

func main() {

	router := gin.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("shop")
	productsCollection = db.Collection("products")
	Salecollection = db.Collection("Sales")

	router.POST("/api/products", Create_product)
	router.POST("/api/sale", Entry_sale)
	router.GET("/Saleview", Agggregate_sale)
	router.Run(":8080")

}
