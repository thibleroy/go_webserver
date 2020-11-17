package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)
func loadMusicRoutes(router *mux.Router) {
	router.HandleFunc("/tracks", controllers.GetTracksController).Methods("GET")
	router.HandleFunc("/tracks", controllers.PostTrackController).Methods("POST")
	router.HandleFunc("/tracks/{id}", controllers.GetTrackController).Methods("GET")

}
