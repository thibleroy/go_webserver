package lib

import "net/http"

type Route struct {
	Name    string
	Handler http.Handler
}
