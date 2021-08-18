package response

import (
	"ticketing/business/movies"
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"original_title"`
	Language    string    `json:"original_language"`
	Description string    `json:"overview"`
	Path        string    `json:"poster_path"`
	VoteAverage float64    `json:"vote_average"`
	VoteCount   int64    `json:"vote_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   time.Time `json:"deleted_at"`
}

func FromDomain(moviesDomain movies.Domain) Movie {
	return Movie{
		ID:          moviesDomain.ID,
		Title:       moviesDomain.Title,
		Language:    moviesDomain.Language,
		Description: moviesDomain.Description,
		Path:        moviesDomain.Path,
		VoteAverage: moviesDomain.VoteAverage,
		VoteCount:   moviesDomain.VoteCount,
		CreatedAt:   moviesDomain.CreatedAt,
		UpdatedAt:   moviesDomain.UpdatedAt,
		// DeletedAt:   moviesDomain.DeletedAt,
	}
}
