package main

import (
	"bytes"
	"fmt"
	"scbam/fcm-publisher/helper"
	"scbam/fcm-publisher/router"
	"scbam/fcm-publisher/usecase"

	"github.com/gin-gonic/gin"
)

type respBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func main() {
	helper.InitConfig()
	port := helper.GetConfig("server.port")

	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()
	router.NewRouter(e)

	decodedKey, err := helper.GetDecodedFireBaseKey()
	if err != nil {
		panic(err)
	}

	// if err := os.Setenv("HTTPS_PROXY", helper.GetConfig("FUNDCLICK.HTTPS_PROXY")); err != nil {
	// 	fmt.Printf("panic Print : " + helper.GetConfig("FUNDCLICK.HTTPS_PROXY") + "\n")
	// 	panic(err)
	// }

	// if err := os.Setenv("HTTP_PROXY", helper.GetConfig("FUNDCLICK.HTTP_PROXY")); err != nil {
	// 	fmt.Printf("panic Print : " + helper.GetConfig("FUNDCLICK.HTTP_PROXY") + "\n")
	// 	panic(err)
	// }

	// proxyStr, userpass, err := helper.GetSystempConfig("google.com")

	if err != nil {
		fmt.Printf("" + err.Error())
		panic(err)
	}

	err = usecase.PushNotification(
		decodedKey,
		"Test",
		"test",
		map[string]string{
			"eggs":    "",
			"bacon":   "",
			"sausage": "",
		},
	)

	if err != nil {
		fmt.Printf("Err")
		panic(err)
	}

	e.Run(":" + port)
}
