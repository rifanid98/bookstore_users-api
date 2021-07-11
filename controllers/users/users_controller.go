package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rifanid98/bookstore_oauth-go/oauth"
	resp "github.com/rifanid98/bookstore_utils-go/response"
)

func init() {
	oauth.BaseUrl = "localhost:8001"
}

func GetById(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, resp.BadRequest("Invalid user id"))
		return
	}

	user, getErr := services.UsersService.GetUserById(userId)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, resp.Success(user.Marshall(false)))
		return
	}

	marshalledUser := user.Marshall(oauth.IsPublic(c.Request))
	c.JSON(http.StatusOK, resp.Success(marshalledUser))
}

func Get(c *gin.Context) {
	query := &users.UserQuery{
		Status: c.Query("status"),
	}

	users, getErr := services.UsersService.GetUser(query)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	marshalledUsers := users.Marshall(oauth.IsPublic(c.Request))
	c.JSON(http.StatusOK, resp.Success(marshalledUsers))
}

func Create(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := resp.BadRequest("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, err := services.UsersService.CreateUser(&u)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	marshalledUser := user.Marshall(c.GetHeader("X-Public") == "true")
	c.JSON(http.StatusCreated, resp.Created(marshalledUser))
	return
}

func Update(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := resp.BadRequest("Invalid user id")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, getErr := services.UsersService.GetUserById(userId)
	if getErr != nil {
		c.JSON(int(getErr.StatusCode), getErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resp.BadRequest("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, err := services.UsersService.UpdateUser(user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	marshalledUser := user.Marshall(c.GetHeader("X-Public") == "true")
	c.JSON(http.StatusOK, resp.Success(marshalledUser))
	return
}

func Delete(c *gin.Context) {
	userId, parseErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := resp.BadRequest("Invalid user id")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user := &users.User{
		Id: userId,
	}

	_, err := services.UsersService.DeleteUser(user)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	c.JSON(http.StatusOK, resp.Success("deleted"))
	return
}

func Login(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := resp.BadRequest("Invalid json input")
		c.JSON(int(restErr.StatusCode), restErr)
		return
	}

	user, err := services.UsersService.LoginUser(&u)
	if err != nil {
		c.JSON(int(err.StatusCode), err)
		return
	}

	marshalledUser := user.Marshall(c.GetHeader("X-Public") == "true")
	c.JSON(http.StatusCreated, resp.Created(marshalledUser))
	return
}
