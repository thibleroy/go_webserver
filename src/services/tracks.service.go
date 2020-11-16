package services
import (
	lib "../lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetTrack(id string) lib.ITrack {
	return lib.ITrack{
		ID:        primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:      "trackTitle",
		Artist:    "artistName",
	}
}
