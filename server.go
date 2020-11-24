package main

import (
    app "go_webserver/src"
    "go_webserver/src/lib"
)

func main(){
    serverEnv := lib.GetServerEnv()
    app.InitServer(serverEnv)
}
