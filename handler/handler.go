package handler

import (
	"fmt"
	"scbam/fcm-publisher/helper"
	"scbam/fcm-publisher/model"
	"scbam/fcm-publisher/usecase"

	"github.com/gin-gonic/gin"
)

func PushNotification(c *gin.Context) {
	decodedKey, err := helper.GetDecodedFireBaseKey()

	if err != nil {
		fmt.Printf("" + err.Error())
		panic(err)
	}

	body := new(model.NotificationItem)

	err = c.BindJSON(body)
	if err != nil {
		fmt.Printf("" + err.Error())
		panic(err)
	}

	err = usecase.PushNotification(
		decodedKey,
		*body,
	)

	if err != nil {
		fmt.Printf("Err")
		panic(err)
	}
}
