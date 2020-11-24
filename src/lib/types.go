package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IResource struct {
	ID primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}
type ITrack struct {
	Resource IResource
	Title string
	Artist string
}
type IEnvironment struct {
	WebServerPort int
	MongoURL string
	MongoPort int
	JwtSecret string
}

type IPostReturn struct {
	ID primitive.ObjectID
}

type IUser struct {
	Resource IResource
	FirstName string
	LastName string
	Email string
	Password string
}

type IError struct {
	Error   error
	Message string
	Code    int
}
