package main

import (
	"github.com/tnadalie/go-app/app"
	"github.com/tnadalie/go-app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
