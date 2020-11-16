package controllers

import (
	"../services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	track := services.GetTrack(id)
	value, _ := json.Marshal(track)
	w.Write(value)
}

