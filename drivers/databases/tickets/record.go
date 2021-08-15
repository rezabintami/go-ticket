package tickets

import (
	"ticketing/business/tickets"
	"time"

	"gorm.io/gorm"
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
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (rec *Tickets) toDomain() tickets.Domain {
	return tickets.Domain{
		ID:          rec.ID,
		BookingCode: rec.BookingCode,
		MovieID:     rec.ID,
		UserID:      rec.UserID,
		TheaterID:   rec.TheaterID,
		Seats:       rec.Seats,
		TotalPrice:  rec.TotalPrice,
		Time:        rec.Time,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		// DeletedAt:   rec.DeletedAt,
	}
}

func fromDomain(ticketsDomain tickets.Domain) *Tickets {
	return &Tickets{
		ID:          ticketsDomain.ID,
		BookingCode: ticketsDomain.BookingCode,
		MovieID:     ticketsDomain.MovieID,
		UserID:      ticketsDomain.UserID,
		TheaterID:   ticketsDomain.TheaterID,
		Seats:       ticketsDomain.Seats,
		TotalPrice:  ticketsDomain.TotalPrice,
		Time:        ticketsDomain.Time,
		CreatedAt:   ticketsDomain.CreatedAt,
		UpdatedAt:   ticketsDomain.UpdatedAt,
		// DeletedAt:   ticketsDomain.DeletedAt,
	}
}
