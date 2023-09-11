package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/megaqstar/web-core/config"
	log "github.com/sirupsen/logrus"
)

var firebaseApp *firebase.App

func connect() error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(fmt.Sprintf("Can't get config: %s", err))
		return err
	}
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: cfg.Json.Firebase.ProjectID}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Error(fmt.Sprintf("Can't connect firebase app: %s", err))
		return err
	}
	firebaseApp = app
	return nil
}

func GetClient() (*firebase.App, error) {
	if firebaseApp == nil {
		err := connect()
		if err != nil {
			return nil, err
		}
	}
	return firebaseApp, nil
}

func Disconnect() {
	firebaseApp = &firebase.App{}
}
