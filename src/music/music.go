package music

import (
	mylog "../middlewares/log"
	"net/http"
	types "../lib"
)

var routePath = "/music"

func musicHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			resp := "<h1> Hello result </h1>"
			w.Write([]byte(resp))
			mylog.MonLog("GET music")
		} else {
			mylog.MonLog("pas GET" + req.Method)
		}
	}
	return http.HandlerFunc(fn)
}

var MusicRoute = types.Route{
	Name: routePath,
	Handler: musicHandler(),
}
