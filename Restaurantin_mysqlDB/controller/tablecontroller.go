package controller

import (
	"net/http"
	model "restaurant_mysql_db/Model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Tablecon struct {
	tablemanage model.TableProcess
	validate    *validator.Validate
}

func Table_controller(tablemanage model.TableProcess) *Tablecon {

	return &Tablecon{tablemanage: tablemanage,
		validate: validator.New(),
	}
}

func (a *Tablecon) ViewTables(c *gin.Context) {

	datas, err := a.tablemanage.Get_All()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to fetch the data"})
		return
	}

	c.JSON(http.StatusOK, datas)
}

func (a *Tablecon) Maketable(c *gin.Context) {

	var temp model.Table

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	exist, _ := a.tablemanage.Viewtable(temp.Table_number)

	if exist.Table_number == temp.Table_number {

		c.JSON(http.StatusOK, gin.H{"Message": "This table_number has been exist"})
		return
	}

	valerr := a.validate.Struct(temp)

	if valerr != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Validation_Error": valerr.Error()})
		return
	}
	data, err := a.tablemanage.Createtable(temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (a *Tablecon) Update_table(c *gin.Context) {

	var temp model.Table

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	valerr := a.validate.Struct(temp)

	if valerr != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Validation_Error": valerr.Error()})
		return
	}

	existingtable, _ := a.tablemanage.Viewtable(temp.Table_id)

	existingtable.Number_of_guest = temp.Number_of_guest
	existingtable.Table_number = temp.Table_number

	updated, err := a.tablemanage.Update_table(existingtable)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Unabletoupdate": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (a *Tablecon) Remove_table(c *gin.Context) {

	table_id := c.Param("table_id")
	removed, err := a.tablemanage.Delete_table(table_id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to remove"})
		return
	}
	c.JSON(http.StatusOK, removed)
}
