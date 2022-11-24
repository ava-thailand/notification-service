package usecase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"google.golang.org/api/option"
)

func PushNotification(
	decodedKey []byte,
	title string,
	body string,
	payload map[string]string,
) (err error) {

	var deviceTokens []string

	deviceTokens = append(
		deviceTokens,
		"cQ5QVImXSMutipD2SpOdKM:APA91bE0hsE2PsDY0-iesww5kEmXzQtZT1GDTLU_hXVPTXui31YNBdXkU-eiw0dn80fLtQemzsgyhl6P-Z9rn8rlM4nxsWyVI1Ue6ThNOeYh3kIee4cMLEpjli_cejQfKOq3f7QxNl2u",
	)

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

	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data:   payload,
		Tokens: deviceTokens, // it's an array of device tokens
	})

	if err != nil {
		panic(err)
	}

	print("\nResponse success count : ", response.SuccessCount)
	print("\nResponse failure count : ", response.FailureCount)
	print("\n")

	return
}
