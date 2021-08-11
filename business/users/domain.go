package users

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	Name       string
	Password   string
	Email      string
	Balance    float64
	Language   string
	Created_At time.Time
	Updated_At time.Time
	Token      string
}

type UseCase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}

type Repository interface {
	Login(ctx context.Context, email, password string) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}
