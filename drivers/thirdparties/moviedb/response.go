package moviedb

import (
	"ticketing/business/moviedb"
)

// type Response struct {
// 	Results []Result `json:"results"`
// }

// type Result struct {
// 	ID          int    `json:"id"`
// 	Title       string `json:"original_title"`
// 	Language    string `json:"original_language"`
// 	Description string `json:"overview"`
// 	Path        string `json:"poster_path"`
// 	VoteAverage string `json:"vote_average"`
// 	VoteCount   string `json:"vote_count"`
// }

type Response struct {
	Result []struct {
		ID          int64   `json:"id"`
		Title       string  `json:"original_title"`
		Language    string  `json:"original_language"`
		Description string  `json:"overview"`
		Path        string  `json:"poster_path"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
	} `json:"results"`
}

func toDomain(resp Response) []moviedb.Domain {
	movies := []moviedb.Domain{}
	for _, value := range resp.Result {
		mov := moviedb.Domain{
			MovieID:     value.ID,
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
