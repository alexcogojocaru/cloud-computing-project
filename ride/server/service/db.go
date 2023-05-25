package service

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreWrapper struct {
	Client *firestore.Client
}

func NewFirestoreDb() *FirestoreWrapper {
	conf := &firebase.Config{ProjectID: "cloudcomputing-386413"}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		return nil
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil
	}

	return &FirestoreWrapper{
		Client: client,
	}
}
