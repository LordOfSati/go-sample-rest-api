package controllers

import (
	"fmt"
	"net/http"
)

// GetAllMovies ...
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
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
