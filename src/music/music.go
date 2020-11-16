package music

import (
	mylog "../middlewares/log"
	"net/http"
	types "../lib"
 	"encoding/json"
)

var routePath = "/music"



func musicHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			resp, _ := json.Marshal(types.IMusic{Name: "musicName", Artist: "artistName"})
			w.Write(resp)
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
