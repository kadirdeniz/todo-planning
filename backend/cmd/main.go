package main

import (
	"go.uber.org/zap"
)

func main() {
	app, err := InitializeApplication()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		app.Logger.Error("Application failed to start", zap.Error(err))
		panic(err)
	}
}