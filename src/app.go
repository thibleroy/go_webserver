package src

import (
	"./music"
	"encoding/json"
	"net/http"
	"fmt"
	types "./lib"
	"strconv"
)

var apiRoutes = []types.Route{music.MusicRoute}


func makeRoutes(routes []types.Route) {
	for _, fun := range routes {
		http.Handle(fun.Name, fun.Handler)
		fmt.Println("Route OK", fun.Name)
	}
}

func InitServer(env types.IEnvironment)  {
	webenv, _ := json.Marshal(env)
	fmt.Println("start server with env", string(webenv))
	makeRoutes(apiRoutes)
	err := http.ListenAndServe(":" + strconv.Itoa(env.Port), nil)
	if err != nil {
		fmt.Println("error", err)
	}
}
