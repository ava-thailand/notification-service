package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"google.golang.org/api/option"

	"scbam/fcm-publisher/model"
)

func PushNotification(
	decodedKey []byte,
	item model.NotificationItem,
) (notificationResponse model.NotificationResponse, err error) {

	opts := []option.ClientOption{
		option.WithCredentialsJSON(decodedKey),
	}

	// Initialize firebase app
	app, err := firebase.NewApp(context.Background(), nil, opts...)

	if err != nil {
		fmt.Printf("err" + err.Error())
		panic(err)
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		fmt.Printf("err" + err.Error())
		panic(err)
	}
	payloadMap := map[string]string{
		"\"id\"":       strconv.Itoa(item.Payload.Id),
		"\"type\"":     "\"" + item.Payload.Type + "\"",
		"\"date\"":     strconv.Itoa(item.Payload.Date),
		"\"title\"":    "\"" + item.Payload.Title + "\"",
		"\"content\"":  "\"" + item.Payload.Content + "\"",
		"\"url\"":      "\"" + item.Payload.Url + "\"",
		"\"fundCode\"": "[\"" + strings.Join(item.Payload.FundCode, "\",\"") + "\"]",
	}

	if len(item.Payload.Image) > 0 {
		payloadMap["image"] = item.Payload.Image
	}

	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: item.Title,
			Body:  item.Body,
		},
		Data:   payloadMap,
		Tokens: item.DeviceTokens, // it's an array of device tokens
	})

	if err != nil {
		fmt.Printf("err" + err.Error())
		panic(err)
	}

	notificationResponse.SuccessCount = response.SuccessCount
	notificationResponse.FailureCount = response.FailureCount
	print("\nResponse success count : ", response.SuccessCount)
	print("\nResponse failure count : ", response.FailureCount)
	print("\n")

	return
}
