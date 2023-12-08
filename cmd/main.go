package main

import (
	"context"
	"log"

	"github.com/lamoda-tech/loikx/internal/app"
)

func main() {
	application := app.NewApp()
	if err := application.Init(context.Background()); err != nil {
		log.Fatalf("%v\n", err)
	}

	if err := application.Start(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
