package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := errors.BadRequestError("Invalid user id")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}
