package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)


func loadMusicRoutes(router *mux.Router) {
	router.HandleFunc("/tracks/{id}", controllers.GetTrackController).Methods("GET")
}
