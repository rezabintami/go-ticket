package tickets

import "time"

type Tickets struct {
	ID          int       `json:"id"`
	BookingCode string    `json:"booking_code"`
	MovieID     int       `json:"movie_id"`
	UserID      string    `json:"user_id"`
	TheaterID   float64   `json:"theater_id"`
	Seats       string    `json:"seats"`
	TotalPrice  int       `json:"total_price"`
	Time        time.Time `json:"time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
