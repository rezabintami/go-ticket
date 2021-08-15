package tickets

import (
	"ticketing/business/tickets"
	"ticketing/drivers/databases/movies"
	"ticketing/drivers/databases/theater"
	"ticketing/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Tickets struct {
	ID          int    `json:"id"`
	BookingCode string `json:"booking_code"`
	MovieID     int    `json:"movie_id"`
	Movie       movies.Movie
	UserID      int `json:"user_id"`
	User        users.Users
	TheaterID   int            `json:"theater_id"`
	Theater     theater.Theater
	Seats       string         `json:"seats"`
	TotalPrice  float64        `json:"total_price"`
	Time        time.Time      `json:"time"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (rec *Tickets) toDomain() tickets.Domain {
	return tickets.Domain{
		ID:          rec.ID,
		BookingCode: rec.BookingCode,
		MovieID:     rec.MovieID,
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
