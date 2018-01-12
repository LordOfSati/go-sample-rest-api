package controller

import (
	"fmt"
	"net/http"
	"strconv"

	repository "github.com/LordOfSati/go-sample-rest-api/movies/repository"
	"github.com/gorilla/mux"
)

var movieRepository = repository.MovieRepository{}

func init() {
	movieRepository.Server = "localhost"
	movieRepository.Database = "movies"
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
	movieID, _ := strconv.Atoi(vars["id"])
	if movie, err := movieRepository.FindByID(movieID); err != nil {
		SendError(w, err)
	} else {
		SendResponse(w, movie)
	}
}

// CreateMovie - To create new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("CreateMovie..")
}

// UpdateMovie - To update an existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("UpdateMovie..")
}

// DeleteMovie - To delete an existing movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("DeleteMovie..")
}
