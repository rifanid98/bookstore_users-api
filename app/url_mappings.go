package app

import (
	ping "bookstore_users-api/controllers/ping"
	users "bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("users", users.CreateUser)
	router.GET("users/", users.GetUser)
	router.GET("users/:user_id", users.GetUserById)
	router.PATCH("users/:user_id", users.UpdateUser)
	router.DELETE("users/:user_id", users.DeleteUser)
}
