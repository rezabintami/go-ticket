package response

import (
	"ticketing/business/tickets"
	"time"
)

type Tickets struct {
	ID          int       `json:"id"`
	BookingCode string    `json:"booking_code"`
	MovieID     int       `json:"movie_id"`
	UserID      int       `json:"user_id"`
	TheaterID   int       `json:"theater_id"`
	Seats       string    `json:"seats"`
	TotalPrice  float64   `json:"total_price"`
	Time        time.Time `json:"time"`
}

func FromDomain(ticketDomain []tickets.Domain) []Tickets {
	tickets := []Tickets{}
	for _, value := range ticketDomain {
		tick := Tickets{
			ID:          value.ID,
			BookingCode: value.BookingCode,
			MovieID:     value.MovieID,
			TheaterID:   value.TheaterID,
			Seats:       value.Seats,
			TotalPrice:  value.TotalPrice, 
			Time: value.Time,
			UserID: value.UserID,
		}
		tickets = append(tickets, tick)
	}
	return tickets
}
