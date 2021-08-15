package request

import (
	"ticketing/business/tickets"
	"time"
)

type Tickets struct {
	MovieID     int       `json:"movie_id"`
	UserID      int       `json:"user_id"`
	TheaterID   int       `json:"theater_id"`
	Seats       string    `json:"seats"`
	TotalPrice  float64   `json:"total_price"`
	Time        time.Time `json:"time"`
}

func (req *Tickets) ToDomain() *tickets.Domain {
	return &tickets.Domain{
		MovieID:     req.MovieID,
		UserID:      req.UserID,
		TheaterID:   req.TheaterID,
		Seats:       req.Seats,
		TotalPrice:  req.TotalPrice,
		Time:        req.Time,
	}
}
