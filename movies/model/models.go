package model

import "gopkg.in/mgo.v2/bson"

// Movie ...
type Movie struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	ReleaseYear int           `json:"releaseYear" bson:"release_year"`
	Length      int           `json:"length" bson:"length"`
	Genre       []string      `json:"genre" bson:"genre"`
}
