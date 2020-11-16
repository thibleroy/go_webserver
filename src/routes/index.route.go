package routes

import (
	"github.com/gorilla/mux"
)

func LoadRouters(router *mux.Router) {
	loadMusicRoutes(router)
}
