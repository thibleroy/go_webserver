package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go_webserver/src/services"
	"net/http"
)

func GetTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for get in controller", id)
	services.RetrieveTrack(id)
}
