package main

import (
    app "./src"
    "./src/lib"
)

func main(){
    serverEnv := lib.GetServerEnv()
    app.InitServer(serverEnv)
}
