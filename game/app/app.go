package app

import (
	"log"
	"net/http"

	"github.com/BernardN38/wasteland-exchange-server/game"
	"github.com/BernardN38/wasteland-exchange-server/router"
	"github.com/BernardN38/wasteland-exchange-server/service"
)

type App struct {
	Router *router.Router
	Config Config
}

func NewApp() *App {
	g, err := game.NewGame()
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}
	r := router.NewRouter(service.NewService(g))
	return &App{Router: r, Config: NewConfig()}
}

func (a *App) Listen() error {
	log.Println("listening")
	log.Fatal(http.ListenAndServe(a.Config.Port, a.Router.Mux))
	return nil
}
