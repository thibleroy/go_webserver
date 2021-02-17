package routes

import (
	"github.com/gorilla/mux"
	"go_webserver/src/controllers"
)
func loadTracksRoutes(router *mux.Router) {
	router.HandleFunc("/tracks/{id}", controllers.GetTrackController).Methods("GET")
}
