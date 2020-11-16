package src

import (
	"./lib"
	"./lib/db"
	"./lib/http"
	"./routes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)


func InitServer(env lib.IEnvironment) {
	APIEnv, _ := json.Marshal(env)
	fmt.Println("API Environment", string(APIEnv))

	lib.Router = mux.NewRouter()
	routes.LoadRouters(lib.Router)

	dbName := "MyMusicAPI"

	lib.MyMusicAPIDB, lib.DBContext = db.InitDB(env.MongoURL, env.MongoPort, dbName)
	http.InitWebServer(env.WebServerPort, lib.Router)
}
