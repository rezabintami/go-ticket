package topup

import (
	"context"
	"ticketing/business/payments"
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
	CreateTransactions(ctx context.Context, data *payments.Domain) (payments.DomainResponse, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}
