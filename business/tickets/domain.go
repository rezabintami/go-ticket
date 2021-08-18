package tickets

import (
	"context"
	"time"
)

type Domain struct {
	ID          int       `json:"id"`
	BookingCode string    `json:"booking_code"`
	MovieID     int       `json:"movie_id"`
	UserID      int       `json:"user_id"`
	TheaterID   int       `json:"theater_id"`
	Seats       string    `json:"seats"`
	TotalPrice  float64   `json:"total_price"`
	Time        time.Time `json:"time"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	// DeletedAt   gorm.DeletedAt `json:"deleted_at"`
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
