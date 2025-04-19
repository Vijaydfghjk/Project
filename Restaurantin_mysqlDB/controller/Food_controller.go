package controller

import (
	"net/http"
	model "restaurant_mysql_db/Model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Foodcon struct {
	foodnmanage model.Foodprocess
	validator   *validator.Validate
}

func Food_controll(foodnmanage model.Foodprocess) *Foodcon {

	return &Foodcon{foodnmanage: foodnmanage,

		validator: validator.New(),
	}
}

func (a *Foodcon) Creatfood(c *gin.Context) {

	var temp model.Food
	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if validation_error := a.validator.Struct(temp); validation_error != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Validation_error": validation_error.Error()})
		return
	}

	inserted, err := a.foodnmanage.Add_food(temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to insertfood"})
		return
	}

	c.JSON(http.StatusOK, inserted)
}

// Getfoods retrieves a list of available food items
// @Summary Get all food items
// @Description Retrieves all food items available in the restaurant
// @Tags Food
// @Produce json
// @Success 200 {array} model.Food
// @Failure 500 {object} map[string]string
// @Router /food [get]

func (a *Foodcon) Getfoods(c *gin.Context) {

	foods, err := a.foodnmanage.Viewfoods()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func (a *Foodcon) Remove_food(c *gin.Context) {

	food_id := c.Param("foodid")

	removed, err := a.foodnmanage.Removefood(food_id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to delete"})
		return
	}
	c.JSON(http.StatusOK, removed)
}
