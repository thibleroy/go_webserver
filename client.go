package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	var resp, err = http.Get("http://localhost:8080/music")
	 if resp != nil {
		 body, _ := ioutil.ReadAll(resp.Body)
		 if resp.StatusCode == 200 {
			 fmt.Println("body value", string(body))
		 } else {
			 fmt.Println("http error", resp.Status)
		 }
	 } else {
		 fmt.Println("Error", err)
	 }
}
