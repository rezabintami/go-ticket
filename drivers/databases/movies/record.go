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
	Description string    `json:"overview"`
	Path        string    `json:"poster_path"`
	VoteAverage float64    `json:"vote_average"`
	VoteCount   int64    `json:"vote_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (rec *Movie) toDomain() movies.Domain {
	return movies.Domain{
		ID:          rec.ID,
		Title:       rec.Title,
		Language:    rec.Language,
		Description: rec.Description,
		Path:        rec.Path,
		VoteAverage: rec.VoteAverage,
		VoteCount:   rec.VoteCount,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		// DeletedAt:   rec.DeletedAt,
	}
}

func fromDomain(movieDomain moviedb.Domain) *Movie {
	return &Movie{
		ID:          movieDomain.ID,
		Title:       movieDomain.Title,
		Language:    movieDomain.Language,
		Description: movieDomain.Description,
		Path:        movieDomain.Path,
		VoteAverage: movieDomain.VoteAverage,
		VoteCount:   movieDomain.VoteCount,
		// CreatedAt:   movieDomain.CreatedAt,
		// UpdatedAt:   movieDomain.UpdatedAt,
		// DeletedAt:   movieDomain.DeletedAt,
	}
}
