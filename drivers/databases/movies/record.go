package movies

import (
	"ticketing/business/moviedb"
	"ticketing/business/movies"
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"original_title"`
	Language    string    `json:"original_language"`
	MovieID     int64     `json:"movie_id"`
	Description string    `json:"overview"`
	Path        string    `json:"poster_path"`
	VoteAverage float64   `json:"vote_average"`
	VoteCount   int64     `json:"vote_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (rec *Movie) toDomain() movies.Domain {
	return movies.Domain{
		ID:          rec.ID,
		Title:       rec.Title,
		Language:    rec.Language,
		MovieID:     rec.MovieID,
		Description: rec.Description,
		Path:        rec.Path,
		VoteAverage: rec.VoteAverage,
		VoteCount:   rec.VoteCount,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(movieDomain moviedb.Domain) *Movie {
	return &Movie{
		ID:          movieDomain.ID,
		Title:       movieDomain.Title,
		Language:    movieDomain.Language,
		MovieID:     movieDomain.MovieID,
		Description: movieDomain.Description,
		Path:        movieDomain.Path,
		VoteAverage: movieDomain.VoteAverage,
		VoteCount:   movieDomain.VoteCount,
	}
}
