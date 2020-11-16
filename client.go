package main

import (
	types "./src/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	var resp, err = http.Get("http://localhost:8080/tracks/123")
	 if resp != nil {
		 body, _ := ioutil.ReadAll(resp.Body)
		 if resp.StatusCode == 200 {
			 fmt.Println("body value", string(body))
			 var myMusic types.ITrack
			 err := json.Unmarshal(body, &myMusic)
			 if err == nil {
				 fmt.Println("music name", myMusic.Title)
				 fmt.Println("artist name", myMusic.Artist)
			 } else {
				 fmt.Println("error", err)
			 }
		 } else {
			 fmt.Println("http error", resp.Status)
		 }
	 } else {
		 fmt.Println("Connexion error", err)
	 }
}
