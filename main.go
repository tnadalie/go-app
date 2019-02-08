package main

import (
	"github.com/tnadalie/app-go/app"
	"github.com/tnadalie/app-go/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
