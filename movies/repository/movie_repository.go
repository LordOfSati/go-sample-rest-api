package repository

import (
	model "github.com/LordOfSati/go-sample-rest-api/movies/model"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MovieRepository ..
type MovieRepository struct {
	Server     string
	Database   string
	Collection string
}

var database *mongo.Database

// Connect - To connect with the mongodb server
func (movieRepository *MovieRepository) Connect() {
	if session, err := mongo.Dial(movieRepository.Server); err != nil {
		panic(err)
	} else {
		database = session.DB(movieRepository.Database)
	}
}

// FindAllMovies - To find all documents from movies collection
func (movieRepository *MovieRepository) FindAllMovies() ([]model.Movie, error) {
	var movies []model.Movie
	err := database.C(movieRepository.Collection).Find(bson.M{}).All(&movies)
	return movies, err
}

// FindByID - To find a movie by ID
func (movieRepository *MovieRepository) FindByID(id string) (model.Movie, error) {
	var movie model.Movie
	err := database.C(movieRepository.Collection).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert - Insert movie document in to the collection
func (movieRepository *MovieRepository) Insert(movie model.Movie) error {
	err := database.C(movieRepository.Collection).Insert(&movie)
	return err
}

// Delete - Delete an existing movie document from the collection
func (movieRepository *MovieRepository) Delete(movie model.Movie) error {
	err := database.C(movieRepository.Collection).Remove(&movie)
	return err
}

// Update - Update an existing movie document in the collection
func (movieRepository *MovieRepository) Update(movie model.Movie) error {
	err := database.C(movieRepository.Collection).UpdateId(movie.ID, &movie)
	return err
}
