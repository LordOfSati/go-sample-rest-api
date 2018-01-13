package main

import (
	"net/http"

	"github.com/LordOfSati/go-sample-rest-api/movies/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.GetMovieByID).Methods("GET")
	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies", controller.DeleteMovie).Methods("DELETE")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic("Error in server start up ..")
	}
}
