package main

import (
    app "./src"
    "fmt"
    "./src/lib"
    "os"
)

func main(){
    err, serverEnv := lib.GetServerEnv()
    if err != nil{
        fmt.Println("error getting server environment")
        os.Exit(2)
    } else {
        app.InitServer(serverEnv)
    }
}
