package controllers

import (
	"../services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"../lib"
)

func GetTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for get", id)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	track := services.GetTrack(objId)
	value, _ := json.Marshal(track)
	w.Write(value)
}

func GetTracksController (w http.ResponseWriter, req *http.Request) {
	tracks := services.GetTracks()
	value, _ := json.Marshal(tracks)
	w.Write(value)
}

func PostTrackController (w http.ResponseWriter, req *http.Request) {
	var track lib.ITrack
	bodyTrack,_ := ioutil.ReadAll(req.Body)

	err := json.Unmarshal(bodyTrack, &track)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	returnId := services.PostTrack(track)
	value, _ := json.Marshal(lib.IPostReturn{ID: returnId})
	w.Write(value)
}

