package controller

import (
	"curd/errorhandler"
	"curd/model"
	"curd/service"
	"curd/validater"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	var user model.User

	binderr := c.ShouldBindJSON(&user)
	if binderr != nil {
		e := errorhandler.GetErorr(binderr, http.StatusNotFound, "Add User Bind issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	validateErr := validater.UserAdd(&user)
	if validateErr != nil {
		e := errorhandler.GetErorr(validateErr, http.StatusNotFound, "Add User validate issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	serviceErr := service.AddUser(&user)
	if serviceErr != nil {
		e := errorhandler.GetErorr(serviceErr, http.StatusNotFound, "Add User servive issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "User added successfully.!"})

}

func UpdateUser(c *gin.Context) {
	var user model.User

	binderr := c.ShouldBindJSON(&user)

	if binderr != nil {
		e := errorhandler.GetErorr(binderr, http.StatusNotFound, "Update User Bind issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	validateErr := validater.UserUpdate(&user)
	if validateErr != nil {
		e := errorhandler.GetErorr(validateErr, http.StatusNotFound, "Update User validate issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	serviceErr := service.UserUpdate(&user)
	if serviceErr != nil {
		e := errorhandler.GetErorr(serviceErr, http.StatusNotFound, "Update User servive issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully.!"})

}

func FindUser(c *gin.Context) {
	var req model.Request
	binderr := c.ShouldBindJSON(&req)
	if binderr != nil {
		e := errorhandler.GetErorr(binderr, http.StatusNotFound, "Find user bind issue")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
	}

	validateErr := validater.UserReq(&req)
	if validateErr != nil {
		e := errorhandler.GetErorr(validateErr, http.StatusNotFound, "Find User validate issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	user, serviceErr := service.FindUser(&req)
	if serviceErr != nil {
		e := errorhandler.GetErorr(serviceErr, http.StatusNotFound, "Find User servive issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

func FindAllUser(c *gin.Context) {

	users, serviceErr := service.FindAllUser()
	if serviceErr != nil {
		e := errorhandler.GetErorr(serviceErr, http.StatusNotFound, "Add All User servive issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeteleUser(c *gin.Context) {

	var req model.Request

	binderr := c.ShouldBindJSON(&req)
	if binderr != nil {
		e := errorhandler.GetErorr(binderr, http.StatusNotFound, "Delete user bind issue")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
	}

	validateErr := validater.UserReq(&req)
	if validateErr != nil {
		e := errorhandler.GetErorr(validateErr, http.StatusNotFound, "Delete User validate issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	serviceErr := service.DeleteUser(&req)
	if serviceErr != nil {
		e := errorhandler.GetErorr(serviceErr, http.StatusNotFound, "Delete User servive issue.")
		c.AbortWithStatusJSON(e.StatusCode, e.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Delete successfully.!"})

}
