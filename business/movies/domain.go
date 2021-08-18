package movies

import (
	"context"
	"ticketing/business/moviedb"
	"time"
)

type Domain struct {
	ID          int       `json:"id"`
	Title       string    `json:"original_title"`
	Language    string    `json:"original_language"`
	MovieID     int64     `json:"movie_id"`
	Description string    `json:"overview"`
	Path        string    `json:"poster_path"`
	VoteAverage float64   `json:"vote_average"`
	VoteCount   int64     `json:"vote_count"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	// DeletedAt   time.Time `json:"deleted_at"`
}

type Usecase interface {
	Fetch(ctx context.Context, urlsearch string, search string) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Check(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *moviedb.Domain) error
	Search(ctx context.Context, title string) ([]Domain, error)
}
