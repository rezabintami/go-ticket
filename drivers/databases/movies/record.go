package movies

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"original_title"`
	Language    string    `json:"original_language"`
	Description string    `json:"overview"`
	Path        float64   `json:"poster_path"`
	VoteAverage string    `json:"vote_average"`
	VoteCount   string    `json:"vote_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
