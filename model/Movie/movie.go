package movie

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"original_title"`
	Language    string    `json:"original_language"`
	Description string    `json:"overview"`
	Path        float64   `json:"poster_path"`
	VoteAverage string    `json:"vote_average"`
	VoteCount   string    `json:"vote_count"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
	Deleted_At  time.Time `json:"deleted_at"`
}
