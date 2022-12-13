package main

import (
	"bytes"
	"scbam/fcm-publisher/helper"
	"scbam/fcm-publisher/router"

	"github.com/gin-gonic/gin"
)

type respBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func main() {
	helper.InitConfig()
	port := helper.GetConfig("server.port")

	gin.SetMode(gin.DebugMode)

	e := gin.Default()
	router.NewRouter(e)

	e.Run(":" + port)
}
