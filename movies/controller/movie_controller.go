package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/satzrp/go-sample-rest-api/movies/model"
	repository "github.com/satzrp/go-sample-rest-api/movies/repository"
	"gopkg.in/mgo.v2/bson"
)

var movieRepository = repository.MovieRepository{}

func init() {
	movieRepository.Server = "localhost"
	movieRepository.Database = "movies-new"
	movieRepository.Collection = "movies"
	movieRepository.Connect()
}

// GetAllMovies - To get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if movies, err := movieRepository.FindAllMovies(); err != nil {
		SendError(w, err)
	} else {
		SendResponse(w, movies)
	}
}

// GetMovieByID - To get a movie by ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	movieID := vars["id"]
	if movie, err := movieRepository.FindByID(movieID); err != nil {
		SendError(w, err)
	} else {
		SendResponse(w, movie)
	}
}

// CreateMovie - To create new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		SendError(w, err)
		return
	}
	movie.ID = bson.NewObjectId()
	if err := movieRepository.Insert(movie); err != nil {
		SendError(w, err)
		return
	}
	SendResponse(w, movie)
}

// UpdateMovie - To update an existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		SendError(w, err)
		return
	}
	if err := movieRepository.Update(movie); err != nil {
		SendError(w, err)
		return
	}
	SendResponse(w, movie)
}

// DeleteMovie - To delete an existing movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		SendError(w, err)
		return
	}
	if err := movieRepository.Delete(movie); err != nil {
		SendError(w, err)
		return
	}
	SendResponse(w, map[string]string{"result": "success"})
}
