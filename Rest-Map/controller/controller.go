package controller

import (
	"net/http"
	"restapi/Model"
	"restapi/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controll struct {
	Service Model.Repository
}

func Mycontroll(service Model.Repository) *Controll {

	return &Controll{Service: service}
}

func (a *Controll) NewUser(c *gin.Context) {

	var temp Model.Register

	err := c.ShouldBind(&temp)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"errr": err.Error()})
		return
	}

	newuser, err := a.Service.Createnewuser(temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newuser)
}
func (a *Controll) Mylogin(c *gin.Context) {

	var temp Model.Login

	err := c.ShouldBind(&temp)

	if err != nil {

		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})

	}

	loguser, err := a.Service.Loginuser(temp)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	token, err := middleware.Generatetoken(int(loguser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
func (a *Controll) Createstudent(c *gin.Context) {

	var student Model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	student.UserID = userID.(uint)

	values, err := a.Service.Createlist(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, values)

}

func (a *Controll) Getstudent(c *gin.Context) {

	userID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	values, err := a.Service.Getall(userID.(uint))

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

	userID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	values, err := a.Service.GetbyID(userID.(uint), id)

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
	userID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}
	students.ID = uint(id)
	students.UserID = userID.(uint)

	updatedstudent, err := a.Service.GetbyID(userID.(uint), id)

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
	updatedone, _ := a.Service.Update(userID.(uint), updatedstudent)
	c.JSON(http.StatusOK, updatedone)
}

func (a *Controll) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("SID"))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found"})
		return
	}

	dell, anerror := a.Service.Delete(userID.(uint), id)

	if anerror != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"errr": err.Error()})
	}

	c.JSON(http.StatusOK, dell)
}
