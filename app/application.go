package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/rifanid98/bookstore_utils-go/env"
	"github.com/rifanid98/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("start the application")
	router.Run(":8000")
}
