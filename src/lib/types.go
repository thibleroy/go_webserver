package lib

import "net/http"

type Route struct {
	Name    string
	Handler http.Handler
}

type IMusic struct {
	Name string
	Artist string
}

type IEnvironment struct {
	Port int
}
