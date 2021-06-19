package app

import (
	controller "bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)
}
