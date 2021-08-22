package tickets

import (
	"context"
	"time"
)

type Domain struct {
	ID          int      
	BookingCode string   
	MovieID     int       
	UserID      int       
	TheaterID   int     
	Seats       string   
	TotalPrice  float64   
	Time        time.Time 
	CreatedAt   time.Time 
	UpdatedAt   time.Time
}

type Usecase interface {
	Delete(ctx context.Context, id int, userId int) error
	GetByID(ctx context.Context, id int) ([]Domain, error)
	Store(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	Delete(ctx context.Context, id int) error
	GetAllByID(ctx context.Context, id int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
