# Simple REST API using GO

This is a simple REST API implementation using [Go](https://golang.org) programming langauge.

## Packages Used

* gorilla/mux - for routing
* gopkg.in/mgo.v2 - for accessing mongodb database

## Endpoints

Base URL: localhost:8080

* GET /movies - To get all the movies from database
* GET /movies/{id} - To get movie by id
* POST /movies - To insert new movie
* PUT /movies - To update an existing movie
* DELETE /movies - To delete a movie