package movies

import (
	"context"
	"ticketing/business/moviedb"
	"time"
)

type Domain struct {
	ID          int
	Title       string
	Language    string
	MovieID     int64
	Description string
	Path        string
	VoteAverage float64
	VoteCount   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
