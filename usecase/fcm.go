package usecase

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"google.golang.org/api/option"

	"scbam/fcm-publisher/model"
)

func PushNotification(
	decodedKey []byte,
	item model.NotificationItem,
) (err error) {

	opts := []option.ClientOption{
		option.WithCredentialsJSON(decodedKey),
	}

	// Initialize firebase app
	app, err := firebase.NewApp(context.Background(), nil, opts...)

	if err != nil {
		panic(err)
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		panic(err)
	}

	var payloadMap map[string]string
	data, err := json.Marshal(item.Payload)
	if err != nil {
		panic(err)
	}

	if json.Unmarshal(data, &payloadMap) != nil {
		panic(err)
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
		panic(err)
	}

	print("\nResponse success count : ", response.SuccessCount)
	print("\nResponse failure count : ", response.FailureCount)
	print("\n")

	return
}
