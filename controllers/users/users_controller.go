package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	resp "bookstore_users-api/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, resp.BadRequest("Invalid user id"))
		return
	}

	user, getErr := services.GetUserById(userId)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	c.JSON(http.StatusOK, resp.Success(user))
}

func GetUser(c *gin.Context) {
	query := &users.UserQuery{
		Status: c.Query("status"),
	}

	users, getErr := services.GetUser(query)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	c.JSON(http.StatusOK, resp.Success(users))
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resp.BadRequest("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	c.JSON(http.StatusCreated, resp.Created(result))
	return
}

func UpdateUser(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := resp.BadRequest("Invalid user id")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, getErr := services.GetUserById(userId)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resp.BadRequest("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	c.JSON(http.StatusOK, resp.Success(result))
	return
}

func DeleteUser(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := resp.BadRequest("Invalid user id")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user := &users.User{
		Id: userId,
	}

	_, err := services.DeleteUser(user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	c.JSON(http.StatusOK, resp.Success("deleted"))
	return
}
