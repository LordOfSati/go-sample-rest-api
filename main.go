package main

import (
	"fmt"
	"net/http"

	"github.com/LordOfSati/go-sample-rest-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	router := mux.NewRouter()
	router.HandleFunc("/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.GetMovieByID).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies", controllers.DeleteMovie).Methods("DELETE")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic("Error in server start up ..")
	}
}
