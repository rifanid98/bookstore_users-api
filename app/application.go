package app

import (
	_ "bookstore_users-api/utils/env"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8000")
}
