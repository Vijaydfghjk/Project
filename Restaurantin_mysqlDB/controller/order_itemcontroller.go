package controller

import (
	"net/http"
	model "restaurant_mysql_db/Model"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type order_itemcon struct {
	ordermange model.Orderprocess
	validate   *validator.Validate
}

type OrderItempack struct {
	Order_items []model.Order_item `json:"order_items"`
}

func OrderItem_controller(ordermange model.Orderprocess) *order_itemcon {

	return &order_itemcon{ordermange: ordermange,

		validate: validator.New(),
	}
}

func (a *order_itemcon) Creater_orderItem(c *gin.Context) {

	var temp OrderItempack

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	var orders []model.Order_item
	var orderchan chan model.Order_item = make(chan model.Order_item, len(temp.Order_items))
	var errchan chan error = make(chan error, len(temp.Order_items))
	var wg sync.WaitGroup
	for _, orderitem := range temp.Order_items {

		wg.Add(1)

		go func(item model.Order_item) {
			defer wg.Done()

			if validation_err := a.validate.Struct(orderitem); validation_err != nil {

				errchan <- validation_err
				return
			}

			orderchan <- orderitem

		}(orderitem)

	}
	wg.Wait()
	close(orderchan)
	close(errchan)

	for err := range errchan {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
	}

	for process := range orderchan {

		orders = append(orders, process)
	}
	inserted, err := a.ordermange.Create_orderItem(orders)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Unable_to_insert": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inserted)
}

func (a *order_itemcon) Update_orderItem(c *gin.Context) {

	var temp model.Order_item

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if validation_err := a.validate.Struct(temp); validation_err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Validation": validation_err.Error()})
		return
	}

	existingorder, _ := a.ordermange.Filter(temp.Order_item_id)

	existingorder.Table_id = temp.Table_id
	existingorder.Quantity = temp.Quantity
	existingorder.Food_id = temp.Food_id

	updatedorder, err := a.ordermange.Update_orderItem(existingorder)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to insert"})
		return
	}

	c.JSON(http.StatusOK, updatedorder)
}
func (a *order_itemcon) Remove_order_item(c *gin.Context) {

	order_id := c.Param("order_item_id")

	removed, err := a.ordermange.Delete_orderItem(order_id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to Remove"})
		return
	}
	c.JSON(http.StatusOK, removed)
}

func (a *order_itemcon) Getall(c *gin.Context) {

	orders, err := a.ordermange.Vieworders()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to fetch the data"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (a *order_itemcon) ViewyTable(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	tableNumber := c.Param("table_number")

	if tableNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table number is required"})
		return
	}

	err := model.CreateOrdersView(db)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orders, err := FetchOrdersByTableNumber(db, tableNumber)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to find"})
	}

	var grandTotal float64

	for _, order := range orders {

		grandTotal += float64(order.Price) * float64(order.Quantity)
	}

	c.JSON(http.StatusOK, gin.H{

		"Orders":       orders,
		"Grand_Total":  grandTotal,
		"Table_number": tableNumber,
	})

}

func FetchOrdersByTableNumber(db *gorm.DB, tableNumber string) ([]model.OrderSummary, error) {

	var orders []model.OrderSummary

	err := db.Raw(`
			
			        select

					foodname,
					price,
					quantity
					from orders
					where table_number = ?
					group by foodname,price, quantity;
			
			
			 `, tableNumber).Scan(&orders).Error

	if err != nil {

		return nil, err
	}
	return orders, nil
}
