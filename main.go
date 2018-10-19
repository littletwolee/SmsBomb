package main

import (
	"SmsBomb/bomb"
	"application"
)

const VERSION = "1.0.1"

var app *application.App

func main() {

	app = application.NewApp("test", "/tmp/config.toml")
	bombModule := bomb.NewSenderModule(app)
	bombModule.Run()
}
