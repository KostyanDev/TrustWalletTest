package main

import (
	"context"
	"log"

	"1/internal/app"
)

func main() {
	appInstance, err := app.InitializeApp()
	if err != nil {
		log.Fatalf("error creating app: %s\n", err)
	}

	if err = appInstance.Run(context.Background()); err != nil {
		log.Fatalf("error running app: %s\n", err)
	}
}
