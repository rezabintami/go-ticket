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
	Description string    `json:"overview"`
	Path        string    `json:"poster_path"`
	VoteAverage float64    `json:"vote_average"`
	VoteCount   int64    `json:"vote_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   time.Time `json:"deleted_at"`
}

// type ThirdParty interface {
// }
// type Usecase interface {
// 	Store(ctx context.Context, search string) error
// }

// type Repository interface {
// 	CheckAndStoreMovies(ctx context.Context, data []moviedb.Domain) error
// }


type Usecase interface {
	Store(ctx context.Context, search string) error
}

type Repository interface {
	Check(ctx context.Context, title string) error
	Store(ctx context.Context, data *moviedb.Domain) error
}
