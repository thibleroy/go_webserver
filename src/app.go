package src

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_webserver/src/lib"
	"go_webserver/src/lib/db"
	"go_webserver/src/lib/http"
	"go_webserver/src/middlewares"
	"go_webserver/src/routes"
)

func InitServer(env lib.IEnvironment) {
	// saves environment in local variable
	lib.Environment = env

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

