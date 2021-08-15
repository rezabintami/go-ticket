package topup

import (
	"context"
	"time"
)

type Domain struct {
	ID        int       `json:"id"`
	User_ID   int       `json:"user_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}
