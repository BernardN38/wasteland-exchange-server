package router

import (
	"net/http"
)

func (r *Router) SetupRoutes() {
	r.Mux.Get("/health", r.CheckHealth)
}
func (router *Router) CheckHealth(w http.ResponseWriter, r *http.Request) {
	result, err := router.Service.GetHealth()
	if err != nil {
		http.Error(w, "Service unhealthy", http.StatusInternalServerError)
	}
	w.Write([]byte(result))
}
