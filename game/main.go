package main

import (
	"github.com/BernardN38/wasteland-exchange-server/app"
)

func main() {
	app := app.NewApp()
	go app.Service.Game.RunGameLoop()
	app.Listen()
}
