package app

import (
	ping "bookstore_users-api/controllers/ping"
	users "bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("users", users.Create)
	router.GET("users/", users.Get)
	router.POST("users/login", users.Login)
	router.GET("users/:user_id", users.GetById)
	router.PATCH("users/:user_id", users.Update)
	router.DELETE("users/:user_id", users.Delete)
}
