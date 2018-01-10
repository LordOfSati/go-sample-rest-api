package controller

import (
	"fmt"
	"net/http"

	repository "github.com/LordOfSati/go-sample-rest-api/movies/repository"
)

var movieRepository = repository.MovieRepository{}

func init() {
	movieRepository.Server = "localhost"
	movieRepository.Database = "movies"
	movieRepository.Collection = "movies"
	movieRepository.Connect()
}

// GetAllMovies ...
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	if movies, err := movieRepository.FindAllMovies(); err != nil {
		fmt.Println("Error in fetching data")
	} else {
		fmt.Println(movies)
	}
	fmt.Println("GetAllMovies..")
}

// GetMovieByID ...
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMoviesById..")
}

// CreateMovie ...
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateMovie..")
}

// UpdateMovie ...
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateMovie..")
}

// DeleteMovie ...
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteMovie..")
}
