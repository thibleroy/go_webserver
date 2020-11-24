package lib

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var MyMusicAPIDB mongo.Database
var DBContext context.Context
var Router *mux.Router
var Environment IEnvironment
