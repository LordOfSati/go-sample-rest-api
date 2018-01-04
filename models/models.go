package models

// Movie ...
type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseYear int    `json:"releaseYear"`
	Length      int    `json:"length"`
}
