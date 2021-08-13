package moviedb

import "context"

type Domain struct {
	ID          int
	Title       string
	Language    string
	Description string
	Path        string
	VoteAverage float64
	VoteCount   int64
}

type Repository interface {
	GetMovies(ctx context.Context, search string) ([]Domain, error)
}
