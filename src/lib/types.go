package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Route struct {
	Name    string
	Handler http.Handler
}

type IResource struct {

}
type ITrack struct{
	ID primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title string
	Artist string
}

type IEnvironment struct {
	WebServerPort int
	MongoURL string
	MongoPort int
}
