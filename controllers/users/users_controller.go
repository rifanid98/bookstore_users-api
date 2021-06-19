package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GetUser Not yet implemented")
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
