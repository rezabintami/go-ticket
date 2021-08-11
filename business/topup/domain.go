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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Usecase interface {
	Payment(ctx context.Context, data *Domain) error
}

type Repository interface {
	Payment(ctx context.Context, data *Domain) error
}
