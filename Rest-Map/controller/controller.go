package controller

import (
	"net/http"
	"restapi/Model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controll struct {
	Service Model.Repository
}

func Mycontroll(service Model.Repository) *Controll {

	return &Controll{Service: service}
}

func (a *Controll) Createstudent(c *gin.Context) {

	var student Model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	values, err := a.Service.Createlist(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, values)

}

func (a *Controll) Getstudent(c *gin.Context) {

	values, err := a.Service.Getall()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, values)
}

func (a *Controll) GetbyID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("SID"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	values, err := a.Service.GetbyID(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, values)
}

func (a *Controll) Updatestudent(c *gin.Context) {

	var students Model.Student
	c.ShouldBind(&students)
	id, err := strconv.Atoi(c.Param("SID"))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing"})
		return

	}

	updatedstudent, err := a.Service.GetbyID(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if students.Name != "" {

		updatedstudent.Name = students.Name
	}
	if students.Place != "" {

		updatedstudent.Place = students.Place
	}

	if students.DOB != "" {

		updatedstudent.DOB = students.DOB
	}

	if students.Contactnumber != "" {

		updatedstudent.Contactnumber = students.Contactnumber
	}
	updatedone, _ := a.Service.Update(updatedstudent)
	c.JSON(http.StatusOK, updatedone)
}

func (a *Controll) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("SID"))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//dellvalue, _ := a.Service.GetbyID(id)

	dell, anerror := a.Service.Delete(id)

	if anerror != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"errr": err.Error()})
	}

	c.JSON(http.StatusOK, dell)
}
