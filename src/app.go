package src

import (
	"./music"
	"net/http"
	"fmt"
	types "./lib"
)

var ApiRoutes = []types.Route{music.MusicRoute}


func makeRoutes(routes []types.Route) {
	for _, fun := range routes {
		http.Handle(fun.Name, fun.Handler)
		fmt.Println("Route OK", fun.Name)
	}
}

func init()  {
	fmt.Println("start server on port")
	makeRoutes(ApiRoutes)
	http.ListenAndServe(":8080", nil)
}
