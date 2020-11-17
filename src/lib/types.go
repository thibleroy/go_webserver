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

type IPostReturn struct {
	ID primitive.ObjectID
}

type HTTPReqInfo struct {
	// GET etc.
	method string
	uri string
	referer string
	ipaddr string
	// response code, like 200, 404
	code int
	// number of bytes of the response sent
	size int64
	// how long did it take to
	duration time.Duration
	userAgent string
}
