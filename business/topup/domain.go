package topup

import (
	"context"
	"ticketing/business/payments"
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	OrderID     string
	PaymentName string
	Name        string
	Amount      float64
	Status      string
	PaymentUrl  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	CreateTransactions(ctx context.Context, payment *payments.Domain, data *Domain, id int) (payments.DomainResponse, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (payments.Domain, error)
	GetByID(ctx context.Context, id int) ([]Domain, error)
}
