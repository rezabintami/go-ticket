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
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
	Deleted_At  time.Time `json:"deleted_at"`
}
