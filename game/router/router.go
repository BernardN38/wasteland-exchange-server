package router

import (
	"github.com/BernardN38/wasteland-exchange-server/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Mux     chi.Router
	Service *service.Service
}

func NewRouter(service *service.Service) *Router {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	r := &Router{Mux: mux, Service: service}
	r.SetupRoutes()
	return r
}
