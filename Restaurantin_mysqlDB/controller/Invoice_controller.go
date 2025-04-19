package controller

import (
	"encoding/json"
	"net/http"
	model "restaurant_mysql_db/Model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Invoice_con struct {
	invoice_handle model.InvoiceProcess
	validation     *validator.Validate
}

func Invoice_controller(invoice_handle model.InvoiceProcess) *Invoice_con {

	return &Invoice_con{invoice_handle: invoice_handle,
		validation: validator.New(),
	}
}
func (a *Invoice_con) Createbill(c *gin.Context) {

	var bill_orders model.Invoice

	db := c.MustGet("db").(*gorm.DB)

	tableDB := model.Table_Repo(db)

	if err := c.ShouldBindJSON(&bill_orders); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	orders, _ := FetchOrdersByTableNumber(db, bill_orders.Table_number)

	jsonData, err := json.Marshal(orders)

	bill_orders.Stored = datatypes.JSON(jsonData)

	inserted, err := a.invoice_handle.CreateInvoice(bill_orders)

	if bill_orders.Payment_status == "PAID" {

		tableDB.Delete_table(bill_orders.Table_number)

	}

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"Invoice": inserted,
	})
}
