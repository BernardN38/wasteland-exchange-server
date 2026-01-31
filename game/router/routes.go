package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (r *Router) SetupRoutes() {
	r.Mux.Get("/api/v1/health", r.CheckHealth)
	r.Mux.Post("/api/v1/game/startnew", r.StartNewGame)
}
func (router *Router) CheckHealth(w http.ResponseWriter, r *http.Request) {
	result, err := router.Service.GetHealth()
	if err != nil {
		http.Error(w, "Service unhealthy", http.StatusInternalServerError)
	}
	w.Write([]byte(result))
}

func (router *Router) StartNewGame(w http.ResponseWriter, r *http.Request) {
	newGame, err := router.Service.StartGame()
	if err != nil {
		http.Error(w, "Failed to start game", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGame)
}

func (router *Router) CreateWebSocket(w http.ResponseWriter, r *http.Request) {

}
