package theater

import (
	"ticketing/business/theater"
	"time"

	"gorm.io/gorm"
)

type Theater struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Place     string         `json:"place"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (rec *Theater) toDomain() theater.Domain {
	return theater.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Place:     rec.Place,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(theaterDomain theater.Domain) *Theater {
	return &Theater{
		ID:        theaterDomain.ID,
		Name:      theaterDomain.Name,
		Place:     theaterDomain.Place,
		CreatedAt: theaterDomain.CreatedAt,
		UpdatedAt: theaterDomain.UpdatedAt,
		DeletedAt: theaterDomain.DeletedAt,
	}
}
