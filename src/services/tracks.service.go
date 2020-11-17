package services

import (
	"../lib"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)


func GetTrack(id primitive.ObjectID) lib.ITrack {
	fmt.Println("objid get", id)
	var trackToRetrieve lib.ITrack
	err := lib.MyMusicAPIDB.Collection("track").FindOne(context.TODO(), bson.M{"id": id}).Decode(&trackToRetrieve)
	if err != nil {
		log.Fatal(err)
	}
	return trackToRetrieve
}

func PostTrack(track lib.ITrack) primitive.ObjectID {
	creationTime := time.Now()
	newTrackId := primitive.NewObjectIDFromTimestamp(creationTime)
	trackToInsert := lib.ITrack{
		ID: newTrackId,
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
		Title:     track.Title,
		Artist:    track.Artist,
	}
	fmt.Println("service post track", trackToInsert)
	_, err := lib.MyMusicAPIDB.Collection("track").InsertOne(context.TODO(), trackToInsert)
	if err != nil {
		log.Fatal("insert error", err)
	}
	return newTrackId
}

func GetTracks() []lib.ITrack {
	var retrievedTracks []lib.ITrack
	cursor,_ :=lib.MyMusicAPIDB.Collection("track").Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedTracks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tracks", retrievedTracks)
	return retrievedTracks
}
