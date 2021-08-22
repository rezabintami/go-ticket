package response

import (
	"ticketing/business/movies"
)

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"original_title"`
	Language    string  `json:"original_language"`
	Description string  `json:"overview"`
	Path        string  `json:"poster_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
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
	}
}

func FromAllDomain(moviesDomain []movies.Domain) []Movie {
	movies := []Movie{}
	for _, value := range moviesDomain {
		mov := Movie{
			ID:          value.ID,
			Title:       value.Title,
			Language:    value.Language,
			Description: value.Description,
			Path:        value.Path,
			VoteAverage: value.VoteAverage,
			VoteCount:   value.VoteCount,
		}
		movies = append(movies, mov)
	}
	return movies
}
