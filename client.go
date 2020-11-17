package main

import (
	"./src/lib"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
)

var url = "http://localhost:8080"

func getRequest(url string){
	var resp, err = http.Get(url)
	if resp != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			fmt.Println("body value", string(body))
		} else {
			fmt.Println("http error", resp.Status)
		}
	} else {
		fmt.Println("Connexion error", err)
	}
}

func postRequest() primitive.ObjectID{
	track := lib.ITrack{
		Title:     "myCreatedTitle",
		Artist:    "myArtist",
	}
	str, _ := json.Marshal(track)
	req, _ := http.NewRequest("POST", url + "/tracks", bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	createdId := resp.Header
	fmt.Println("header loc", createdId)
	return primitive.ObjectID{}
}

func main(){
	id := postRequest()
	fmt.Println("created id", id)
}
