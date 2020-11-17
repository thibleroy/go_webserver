package middlewares

import "github.com/gorilla/mux"

func LoadMiddlewares(router *mux.Router) {
	router.Use(loggingMiddleware)
	router.Use(headersMiddleware)
}