package app

import (
	_ "bookstore_users-api/utils/env"
	"bookstore_users-api/utils/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("start the application")
	router.Run(":8000")
}
