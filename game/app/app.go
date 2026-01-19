package app

import (
	"log"
	"net/http"

	"github.com/BernardN38/wasteland-exchange-server/router"
	"github.com/BernardN38/wasteland-exchange-server/service"
)

type App struct {
	Router *router.Router
	Config Config
}

func NewApp() *App {
	r := router.NewRouter(service.NewService())
	return &App{Router: r, Config: NewConfig()}
}

func (a *App) Listen() error {
	log.Println("listening")
	log.Fatal(http.ListenAndServe(a.Config.Port, a.Router.Mux))
	return nil
}
