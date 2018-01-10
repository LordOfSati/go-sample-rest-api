package model

// Movie ...
type Movie struct {
	ID          int      `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	ReleaseYear int      `json:"releaseYear" bson:"release_year"`
	Length      int      `json:"length" bson:"length"`
	Genre       []string `json:"genre" bson:"genre"`
}
