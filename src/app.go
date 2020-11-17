package src

import (
	"./lib"
	"./lib/db"
	"./lib/http"
	"./middlewares"
	"./routes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func InitServer(env lib.IEnvironment) {
	// prints running environment
	APIEnv, _ := json.Marshal(env)
	fmt.Println("API Environment", string(APIEnv))

	// creates router and loads routes and methods handlers
	lib.Router = mux.NewRouter()
	routes.LoadRouters(lib.Router)

	// loads middlewares
	middlewares.LoadMiddlewares(lib.Router)
	dbName := "MyMusicAPI"

	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB, lib.DBContext = db.InitDB(env.MongoURL, env.MongoPort, dbName)

	// initializes the http server with previously created Mux router
	http.InitWebServer(env.WebServerPort, lib.Router)
}
