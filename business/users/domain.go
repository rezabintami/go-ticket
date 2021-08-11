package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Token     string    `json:"token"`
}

type UserDomain struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	Language string  `json:"language"`
}

type UseCase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (UserDomain, error)
	UpdateUser(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	Login(ctx context.Context, id int) (Domain, error)
	GetByID(ctx context.Context, id int) (UserDomain, error)
	UpdateUser(ctx context.Context, data *Domain, id int) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}
