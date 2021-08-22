package topup

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	Name      string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}
