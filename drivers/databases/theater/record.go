package theater

import "time"

type Theater struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Place     string    `json:"place"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
