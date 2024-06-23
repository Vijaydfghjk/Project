package controller

import (
	"RESTAPi/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Use http.StatusBadRequest (400) for client-side errors.
// Use http.StatusInternalServerError (500) for server-side errors.
type Repo struct {
	value model.Process
}

func Newcontroller(value model.Process) *Repo {

	return &Repo{value: value}
}

func (r *Repo) Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registeredUser, err := r.value.RegisterUser(user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registeredUser)
}

func (r *Repo) Login(c *gin.Context) {

	var logindetails model.Login

	if err := c.ShouldBind(&logindetails); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.value.LoginUser(logindetails.Email, logindetails.Password)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := generateJWT(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (r *Repo) Accountcreation(c *gin.Context) {

	var temp model.Account

	err := c.ShouldBind(&temp)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := r.value.Createaccount(temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (r *Repo) View(c *gin.Context) {

	account, err := r.value.Viewaccount()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (r *Repo) Getbyname(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		log.Fatalf("invalid value")
	}

	account, err := r.value.Getaccount(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (r *Repo) Updateaccount(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var temp model.Account

	if err := c.ShouldBind(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, _ := r.value.Getaccount(id)

	if temp.Name != "" {

		account.Name = temp.Name
	}

	if temp.Adharr != "" {

		account.Adharr = temp.Adharr
	}

	if temp.Email != "" {

		account.Email = temp.Email
	}

	updatedaccount, err := r.value.Updateaccount(account)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedaccount)
}

func (r *Repo) Deleteaccount(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		log.Fatalf("invalid value")
	}

	deleted, err := r.value.Deleteaccount(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deleted)

}
