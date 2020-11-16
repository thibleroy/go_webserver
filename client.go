package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	types "./src/lib"
)

func main(){
	var resp, err = http.Get("http://localhost:8080/music")
	 if resp != nil {
		 body, _ := ioutil.ReadAll(resp.Body)
		 if resp.StatusCode == 200 {
			 fmt.Println("body value", string(body))
			 var myMusic types.IMusic
			 err := json.Unmarshal(body, &myMusic)
			 if err == nil {
				 fmt.Println("music name", myMusic.Name)
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
