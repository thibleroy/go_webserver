package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go_webserver/src/lib"
	"time"
)

const trackCollectionName = "track"

func RetrieveTrack(id primitive.ObjectID) (*lib.ITrack, error) {
	trackToRetrieve := lib.ITrack{}
	err := lib.MyMusicAPIDB.Collection(trackCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&trackToRetrieve)
	if err != nil {
		print("err", err.Error())
		return nil, err
	}
	return &trackToRetrieve, nil
}

func AddTrack(track lib.ITrack) (*primitive.ObjectID, error) {
	fmt.Println("obj", track)
	track.Resource = lib.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(trackCollectionName).InsertOne(context.TODO(), track)
	if err != nil {
		return nil, err
	}
	return &track.Resource.ID, nil
}

func UpdateTrack(track lib.ITrack) (*primitive.ObjectID, error) {
	updateTime := time.Now()
	track.Resource.UpdatedAt = updateTime
	update := bson.M{
		"$set": track,
	}
	_, err := lib.MyMusicAPIDB.Collection(trackCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": track.Resource.ID}, update)
	if err != nil {
		return nil, err
	}
	return &track.Resource.ID, nil
}

func RetrieveAllTracks() (*[]lib.ITrack, error) {
	var retrievedTracks []lib.ITrack
	cursor,_ :=lib.MyMusicAPIDB.Collection(trackCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedTracks)
	if err != nil {
		return nil, err
	}
	return &retrievedTracks, nil
}

func RemoveTrack(id primitive.ObjectID) (*primitive.ObjectID, error){
	_, err := lib.MyMusicAPIDB.Collection(trackCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return nil, err
	}
	return &id, nil
}
