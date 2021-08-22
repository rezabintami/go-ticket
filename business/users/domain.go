package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int     
	Name      string    
	Password  string  
	Email     string    
	Balance   float64  
	Language  string   
	CreatedAt time.Time 
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	UpdateUser(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	// Login(ctx context.Context, id int) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	UpdateUser(ctx context.Context, data *Domain, id int) error
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}
